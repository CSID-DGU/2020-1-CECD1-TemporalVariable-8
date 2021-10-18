package mig

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/Rhymond/go-money"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"lightServer/ent"
	"lightServer/ent/authorize"
	"lightServer/ent/file"
	"lightServer/ent/menu"
	"lightServer/ent/order"
	"lightServer/ent/predicate"
	"lightServer/ent/restaurant"
	"lightServer/ent/user"
	"lightServer/mig/param"
	"lightServer/mig/zaputil"
	"lightServer/pb"
	"math/big"
	"math/rand"
	"mime"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"
)

const (
	CacheBlockchainUrls = "cache.blockchain.urls"
)

func (s *MakItGo) buildGin() (err error) {
	// 로깅 미들웨어
	s.Unsafe.Engine.Use(func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		c.Next()
		//
		end := time.Now()
		latency := end.Sub(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()
		bodySize := c.Writer.Size()
		if raw != "" {
			path = path + "?" + raw
		}
		if len(errorMessage) > 0 {
			s.Unsafe.Logger.Error(
				"[ gin ]",
				zap.Int("status", statusCode),
				zap.Duration("latency", latency),
				zap.String("method", method),
				zap.String("path", path),
				zap.String("client", clientIP),
				zap.String("errors", errorMessage),
			)
		} else {
			s.Unsafe.Logger.Info(
				"[ gin ]",
				zap.Int("status", statusCode),
				zap.Duration("latency", latency),
				zap.String("method", method),
				zap.String("path", path),
				zap.String("client", clientIP),
				zap.Int("body_size", bodySize),
			)
		}
	})
	s.Unsafe.Engine.Use(gin.Recovery())
	// auth
	// /auth/self
	s.Unsafe.Engine.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/static/index.html")
	})

	s.Unsafe.Engine.Static("/static", "./MakItGo/static")
	s.Unsafe.Engine.Static("/files", "./MakItGo/static/files")
	s.Unsafe.Engine.GET("/auth/self", noStore, func(c *gin.Context) {
		var (
			cState, err = c.Cookie("state")
		)
		if err != nil {
			c.JSON(http.StatusOK, nil)
			return
		}
		var (
			session = loadSession(s.Unsafe.Cache, context.Background(), cState)
		)
		c.JSON(http.StatusOK, session)
	})
	// auth
	// /auth/login/:target
	// 동작 : :target을 통해 OAuth2 로그인을 시도함
	s.Unsafe.Engine.GET("/auth/login/:authorizer", noStore, func(c *gin.Context) {
		var (
			oauth2config *oauth2Config
			ok           bool
			state        = randomStateString()
		)
		if oauth2config, ok = s.Configs.oauth2[c.Param("authorizer")]; !ok {
			// TODO
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		//
		c.Redirect(http.StatusTemporaryRedirect, oauth2config.conf.AuthCodeURL(state))
	})
	s.Unsafe.Engine.GET("/auth/logout", noStore, func(c *gin.Context) {
		ck, err := c.Cookie("state")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		_, err = s.Unsafe.Cache.Del(context.Background(), ck).Result()
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		//
		c.SetCookie("state", "<deleted>", -1, "/", "", false, true)
		c.Redirect(http.StatusPermanentRedirect, "/")
	})
	// auth
	// /auth/callback/:target
	// 동작 : :target을 통해 OAuth2 로그인을 시도한 후 그 결과를 받아 옴
	s.Unsafe.Engine.GET("/auth/callback/:target", noStore, func(c *gin.Context) {
		var (
			qState = c.Query("state")
			qCode  = c.Query("code")
		)
		var (
			oauth2config *oauth2Config
			ok           bool
			session      = new(Session)
		)
		if oauth2config, ok = s.Configs.oauth2[c.Param("target")]; !ok {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		tk, err := oauth2config.conf.Exchange(context.Background(), qCode)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		//
		profile, err := oauth2config.getProfile(oauth2config.conf.Client(context.Background(), tk))
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		//
		u, err := s.Unsafe.DB.User.Query().Where(
			user.HasAuthsWith(
				authorize.ProviderEQ(authorize.ProviderGoogle),
				authorize.ServiceID(profile.Get("email").String()),
			)).Only(context.Background())
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		session.Client = u
		//
		err = saveSession(s.Unsafe.Cache, context.Background(), qState, session, s.Configs.DefaultSessionTimeout)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		//
		c.SetCookie("state", qState, int(s.Configs.DefaultSessionTimeout.Seconds()), "/", "", false, true)
		c.Redirect(http.StatusPermanentRedirect, "/")
	})
	// info group
	s.Unsafe.Engine.Group("/info").
		GET("/name", func(c *gin.Context) { c.String(200, s.Configs.Name) })
	// blockchain
	s.Unsafe.Engine.GET("/blockchain/urls", func(c *gin.Context) {
		var urls []string
		err := s.Unsafe.Cache.Get(context.Background(), CacheBlockchainUrls).Scan(&urls)
		if err != nil {
			count, err := s.Unsafe.Contract.ListURLCount(&bind.CallOpts{})
			if err != nil {
				s.Unsafe.Logger.Error("[blockchain]", zap.String("error", err.Error()))
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			var urls = make([]string, count)
			for i := range urls {
				tmp, err := s.Unsafe.Contract.Urls(&bind.CallOpts{}, big.NewInt(int64(i)))
				if err != nil {
					s.Unsafe.Logger.Error("[blockchain]", zap.String("error", err.Error()))
					c.AbortWithStatus(http.StatusInternalServerError)
					return
				}
				urls[i] = tmp
			}
			seterr := s.Unsafe.Cache.Set(context.Background(), CacheBlockchainUrls, urls, time.Minute).Err()
			if seterr != nil {
				s.Unsafe.Logger.Error("[blockchain]", zap.String("error", seterr.Error()), zap.String("type", "cache set error"))
			}
		}
		c.JSON(http.StatusOK, urls)
	})
	s.Unsafe.Engine.GET("/blockchain/address", func(c *gin.Context) {
		c.JSON(http.StatusOK, s.Unsafe.ContractAddress.Hex())
	})
	// File
	// /file/:fid?mime=마임&mode=모드
	// 마임 : <type>/subtype
	// 모드 : view | download | default = view
	s.Unsafe.Engine.GET("/file/:fid", func(c *gin.Context) {
		var (
			pFid  = c.Param("fid")
			qMode = c.Query("mode")
			qMime = c.Query("mime")
		)
		var (
			selectors []predicate.File
		)
		// 파일 uuid 얻어옴
		fuuid, err := uuid.Parse(pFid)
		if err != nil {
			s.Unsafe.Logger.Warn(
				"[ File ]",
				zap.String("message", "Not Valid UUID Format"),
				zap.Namespace("request"),
				zap.String("method", c.Request.Method),
				zap.Stringer("url", c.Request.URL),
				zap.String("proto", c.Request.Proto),
				zap.Object("header", zaputil.ZapHeader(c.Request.Header)),
			)
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		selectors = append(selectors, file.FileID(fuuid))
		// 유저 mime 정보 획득
		if len(qMime) != 0 {
			selectors = append(selectors, file.Mime(qMime))
		}
		// DB 엔티티 획득
		dbfiles, err := s.Unsafe.DB.File.Query().Where(selectors...).All(context.Background())
		if err != nil {
			s.Unsafe.Logger.Warn(
				"[ File ]",
				zap.String("message", "Not Found DB File"),
				zap.Namespace("request"),
				zap.String("method", c.Request.Method),
				zap.Stringer("url", c.Request.URL),
				zap.String("proto", c.Request.Proto),
				zap.Object("header", zaputil.ZapHeader(c.Request.Header)),
			)
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		if len(dbfiles) == 0 {
			s.Unsafe.Logger.Warn(
				"[ File ]",
				zap.String("message", "No DB File"),
				zap.Namespace("request"),
				zap.String("method", c.Request.Method),
				zap.Stringer("url", c.Request.URL),
				zap.String("proto", c.Request.Proto),
				zap.Object("header", zaputil.ZapHeader(c.Request.Header)),
			)
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		//
		var dbfile *ent.File
		switch len(dbfiles) {
		case 0:
			s.Unsafe.Logger.Warn(
				"[ File ]",
				zap.String("message", "Not found db file"),
				zap.Namespace("request"),
				zap.String("method", c.Request.Method),
				zap.Stringer("url", c.Request.URL),
				zap.String("proto", c.Request.Proto),
				zap.Object("header", zaputil.ZapHeader(c.Request.Header)),
			)
			c.AbortWithStatus(http.StatusNotFound)
			return
		default:
			dbfile = dbfiles[0]
		}
		// 실제 파일 얻어옴
		f, fileerr := s.Unsafe.FS.Open(path.Join(fuuid.String(), MimeFile(dbfile.Mime)))
		if fileerr != nil {
			s.Unsafe.Logger.Warn(
				"[ File ]",
				zap.String("message", "Not Found File"),
				zap.Namespace("request"),
				zap.String("method", c.Request.Method),
				zap.Stringer("url", c.Request.URL),
				zap.String("proto", c.Request.Proto),
				zap.Object("header", zaputil.ZapHeader(c.Request.Header)),
			)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		defer f.Close()
		//
		fstat, staterr := f.Stat()
		if staterr != nil {
			s.Unsafe.Logger.Warn(
				"[ File ]",
				zap.String("message", "Not Found File"),
				zap.Namespace("request"),
				zap.String("method", c.Request.Method),
				zap.Stringer("url", c.Request.URL),
				zap.String("proto", c.Request.Proto),
				zap.Object("header", zaputil.ZapHeader(c.Request.Header)),
			)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		//
		switch qMode {
		case "download":
			var ext string
			exts, _ := mime.ExtensionsByType(dbfile.Mime)
			if len(exts) > 0 {
				ext = exts[0]
			}
			c.DataFromReader(http.StatusOK, fstat.Size(), dbfile.Mime, f, map[string]string{
				"Content-Disposition": fmt.Sprintf(`attachment; filename="%s.%s"`, dbfile.Name, ext),
			})
		default:
			c.DataFromReader(http.StatusOK, fstat.Size(), dbfile.Mime, f, nil)
		}
	})
	// RESTful API
	restful := s.Unsafe.Engine.Group("/api", func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	})
	// RESTful API : restaurant
	// /api/restaurants/
	// 동작 : 레스토랑 목록에 대해 전달
	restful.GET("/restaurants", func(ctx *gin.Context) {
		timedCtx, _ := context.WithTimeout(context.Background(), s.Configs.DefaultDBTimeout)
		all := s.Unsafe.DB.Restaurant.Query().WithParent().AllX(timedCtx)
		var names = make([]*pb.RestaurantName, len(all))
		for i, r := range all {
			names[i] = pathTopbrestpath(getRecursiveName(r), r.URI)
		}
		renderByTyped(ctx, http.StatusOK, &pb.RestaurantStatistics{
			Names: names,
		})
	})
	restful.POST("/orders/:order_id/cooking", func(c *gin.Context) {
		var (
			i, _ = strconv.ParseInt(c.Param("order_id"), 10, 64)
		)
		state, _ := c.Cookie("state")
		sess := loadSession(s.Unsafe.Cache, context.Background(), state)
		ord, err := s.Unsafe.DB.Order.Get(context.Background(), int(i))
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		rest, err := s.Unsafe.DB.Order.QueryWhere(ord).WithOwner().Where(restaurant.HasOwnerWith(user.ID(sess.Client.ID))).Only(context.Background())
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		if rest.Edges.Owner.ID != sess.Client.ID {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		cat := time.Now()
		ord.Update().SetCookingAt(cat).SaveX(context.Background())
		renderByTyped(c, http.StatusOK, &pb.ManageCookingResult{
			OrderId: int64(ord.ID),
			TimeAt:  cat.Format(time.RFC3339),
		})
	})
	restful.POST("/orders/:order_id/deliver", func(c *gin.Context) {
		var (
			i, _ = strconv.ParseInt(c.Param("order_id"), 10, 64)
		)
		state, _ := c.Cookie("state")
		sess := loadSession(s.Unsafe.Cache, context.Background(), state)
		ord, err := s.Unsafe.DB.Order.Get(context.Background(), int(i))
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		rest, err := s.Unsafe.DB.Order.QueryWhere(ord).WithOwner().Where(restaurant.HasOwnerWith(user.ID(sess.Client.ID))).Only(context.Background())
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		if rest.Edges.Owner.ID != sess.Client.ID {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		cat := time.Now()
		ord.Update().SetDeliverAt(cat).SaveX(context.Background())
		renderByTyped(c, http.StatusOK, &pb.ManageCookingResult{
			OrderId: int64(ord.ID),
			TimeAt:  cat.Format(time.RFC3339),
		})
	})
	restful.POST("/orders/:order_id/complete", func(c *gin.Context) {
		var (
			i, _ = strconv.ParseInt(c.Param("order_id"), 10, 64)
		)
		state, _ := c.Cookie("state")
		sess := loadSession(s.Unsafe.Cache, context.Background(), state)
		ord, err := s.Unsafe.DB.Order.Get(context.Background(), int(i))
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		rest, err := s.Unsafe.DB.Order.QueryWhere(ord).WithOwner().Where(restaurant.HasOwnerWith(user.ID(sess.Client.ID))).Only(context.Background())
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		if rest.Edges.Owner.ID != sess.Client.ID {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		cat := time.Now()
		ord.Update().SetCompleteAt(cat).SaveX(context.Background())
		renderByTyped(c, http.StatusOK, &pb.ManageCookingResult{
			OrderId: int64(ord.ID),
			TimeAt:  cat.Format(time.RFC3339),
		})
	})
	// RESTful API : restaurant
	// <공통>
	// /api/restaurants/:name[/:subnames]/*?type=[타입]
	// 타입 : json | protobuf | default=json, 출력하는 데이터의 포멧 형식
	//
	// GET /api/restaurants/:name[/:subnames]/
	// 동작 : 음식점의 정보를 가져옴
	// GET /api/restaurants/:name[/:subnames]/menus
	// 동작 : 특정 지점의 메뉴를 모두 가져옴
	// GET /api/restaurants/:name[/:subnames]/order/:order_id
	// 동작 : 특정 주문 정보를 가져옴
	// GET /api/restaurants/:name[/:subnames]/history?start=opt<rfc3339|unixepoch>&end=opt<rfc3339|unixepoch>
	// 동작 : 인증된 범위 내에서 범위내의 주문을 모두 가져옴
	// 예시 01 : 레스토랑 주인 인증됨 : 해당 레스토랑의 범위내 주문 불러 옴
	// 예시 02 : 유저 인증됨 : 해당 레스토랑에서 주문한 범위내 주문 불러 옴
	restful.GET("/restaurants/:name/*lefts", func(c *gin.Context) {
		var (
			name, method, params = restaurantsPathParse(c.Param("name"), c.Param("lefts"))
		)
		r, err := s.Unsafe.DB.Restaurant.Query().
			Where(makeWhere(name)...).WithParent().WithAvatar().Only(context.Background())
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		switch method {
		case "":
			renderByTyped(c, http.StatusOK, dbrestTopbrest(r))
		case "menus":
			menus := r.QueryMenus().Where(menu.IsOption(false)).WithOptions().WithImages().AllX(context.Background())
			renderByTyped(c, http.StatusOK, &pb.Menus{
				Owner: pathTopbrestpath(name, r.URI),
				Menus: dbmenusTopbmenus(menus),
			})
		case "options":
			opts, err := r.QueryMenus().Where(menu.IsOption(true)).All(context.Background())
			if err != nil {
				s.Unsafe.Logger.Error("[req:options]", zap.String("error", err.Error()))
				return
			}
			renderByTyped(c, http.StatusOK, &pb.Menus{
				Owner: pathTopbrestpath(name, r.URI),
				Menus: dbmenusTopbmenus(opts),
			})
		case "category":
			categories := r.QueryCategories().AllX(context.Background())
			renderByTyped(c, http.StatusOK, &pb.Categories{
				Owner:      pathTopbrestpath(name, r.URI),
				Categories: dbcategoriesTopbcategories(categories),
			})
		case "history":
			var (
				start, emptyStart, validStart = timeUnixOrRFC3339(c.Query("start"))
				end, emptyEnd, validEnd       = timeUnixOrRFC3339(c.Query("end"))
			)
			if !validStart || !validEnd {
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}
			q := r.QueryOrders().WithWhere().WithItems().WithWho()
			if !emptyStart {
				q.Where(order.OrderAtGTE(start))
			}
			if !emptyEnd {
				q.Where(order.OrderAtLTE(end))
			}
			orders, err := q.All(context.Background())
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			result := make([]*pb.Order, len(orders))
			for i, o := range orders {
				result[i] = dborderTopborder(name, o)
			}
			renderByTyped(c, http.StatusOK, result)
		case "order":
			if len(params) != 1 {
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}
			var num, err = strconv.ParseInt(params[0], 10, 64)
			if err != nil {
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}
			od, err := r.QueryOrders().Where(order.ID(int(num))).Only(context.Background())
			if err != nil {
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}
			renderByTyped(c, http.StatusOK, dborderTopborder(name, od))
		case "orders":
			ods, err := r.QueryOrders().WithWho().WithItems(func(query *ent.OrderFieldQuery) { query.WithMenu() }).WithWhere().All(context.Background())
			if err != nil {
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}
			renderByTyped(c, http.StatusOK, dbordersTopborders(name, ods))
		}

	})
	restful.POST("/restaurants/:name/*lefts", func(c *gin.Context) {
		var (
			name, method, _ = restaurantsPathParse(c.Param("name"), c.Param("lefts"))
		)
		//state, err := c.Cookie("state")
		//if err != nil {
		//	s.Unsafe.Logger.Error("[req]", zap.String("error", err.Error()))
		//	c.AbortWithStatus(http.StatusUnauthorized)
		//	return
		//}
		//sess := loadSession(s.Unsafe.Cache, context.Background(), state)
		//if sess == nil {
		//	s.Unsafe.Logger.Error("[req]", zap.String("where", "post:session"))
		//	c.AbortWithStatus(http.StatusUnauthorized)
		//	return
		//}
		r, err := s.Unsafe.DB.Restaurant.Query().
			Where(makeWhere(name)...).Only(context.Background())

		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		switch method {
		case "order":
			var p param.ParamOrder
			if err := c.ShouldBindJSON(&p); err != nil {
				//
				s.Unsafe.Logger.Error("[req]", zap.String("err", err.Error()))
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}
			var fields = make([]*ent.OrderField, 0, len(p.Fields))
			for _, field := range p.Fields {
				fields = append(fields,
					s.Unsafe.DB.OrderField.Create().
						SetMenuID(field.Which).
						SetCount(uint16(field.Count)).SaveX(context.Background()),
				)

			}
			ord := s.Unsafe.DB.Order.Create().
				SetWho(s.Configs.guest).
				SetWhere(r).
				AddItems(fields...).
				SaveX(context.Background())

			renderByTyped(c, http.StatusOK, &pb.OrderResult{
				OrderId: int64(ord.ID),
			})
		case "menu":
			var p param.ParamMenu
			if err := c.ShouldBindJSON(&p); err != nil {
				//
				s.Unsafe.Logger.Error("[req]", zap.String("err", err.Error()))
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}
			m, err := s.Unsafe.DB.Menu.Create().
				SetName(p.Name).
				SetDescription(p.Desc).
				SetIsOption(false).
				SetOwner(r).
				SetPrice(money.New(p.Amount, "KRW")).
				AddOptionIDs(p.Options...).Save(context.Background())
			if err != nil {
				s.Unsafe.Logger.Error("[req]", zap.String("error", err.Error()))
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}
			renderByTyped(c, http.StatusOK, dbmenuTopbmenu(m))

		}
	})
	restful.DELETE("/restaurants/:name/*lefts", func(c *gin.Context) {
		var (
			name, method, params = restaurantsPathParse(c.Param("name"), c.Param("lefts"))
		)
		r, err := s.Unsafe.DB.Restaurant.Query().
			Where(makeWhere(name)...).Only(context.Background())
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		fmt.Println(r)
		switch method {
		case "menu":
			if len(params) != 1 {
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}
			target, err := strconv.ParseInt(params[0], 10, 64)
			if err != nil {
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}
			err = s.Unsafe.DB.Menu.DeleteOneID(int(target)).Exec(context.Background())
			if err != nil {
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}
			c.String(http.StatusOK, "")
		}
	})
	return nil
}
func makeWhere(name []string) []predicate.Restaurant {
	switch len(name) {
	case 0:
		return []predicate.Restaurant{}
	case 1:
		return []predicate.Restaurant{restaurant.Name(name[0])}
	default:
		return []predicate.Restaurant{
			restaurant.Name(name[len(name)-1]),
			restaurant.HasParentWith(
				makeWhere(name[:len(name)-1])...,
			),
		}
	}
}
func restaurantsPathParse(start, lefts string) (name []string, method string, param []string) {
	tmp := []string{start}
	for _, s := range strings.Split(lefts, "/") {
		if len(s) > 0 {
			tmp = append(tmp, s)
		}
	}
	//
	if len(tmp) == 1 {
		return tmp, "", nil
	}
	//
	var (
		i int
		e string
	)
	for i, e = range tmp {
		switch e {
		case "menu":
			fallthrough
		case "menus":
			fallthrough
		case "category":
			fallthrough
		case "order":
			fallthrough
		case "orders":
			fallthrough
		case "options":
			fallthrough
		case "history":
			return tmp[:i], e, tmp[i+1:]
		}
	}
	return tmp, method, []string{}

}
func renderByTyped(ctx *gin.Context, code int, obj interface{}) {
	switch ctx.Query("type") {
	default:
		fallthrough
	case "json":
		ctx.JSON(code, obj)
	case "protobuf":
		ctx.ProtoBuf(code, obj)
	}
}

func fileUUIDToURLPath(id uuid.UUID) string {
	return "/file/" + id.String()
}

func dbmenusTopbmenus(menus []*ent.Menu) []*pb.Menu {
	ret := make([]*pb.Menu, len(menus))
	for i, menu := range menus {
		ret[i] = dbmenuTopbmenu(menu)
	}
	return ret
}

func dbmenuTopbmenu(menu *ent.Menu) *pb.Menu {
	result := &pb.Menu{
		Id:          int64(menu.ID),
		Name:        menu.Name,
		Description: menu.Description,
		PriceAmount: menu.Price.Amount(),
		Currency:    menu.Price.Currency().Code,
	}
	result.Options = make([]*pb.Menu, len(menu.Edges.Options))
	for i, mopt := range menu.Edges.Options {
		result.Options[i] = dbmenuTopbmenu(mopt)
	}
	if i, err := menu.Edges.ImagesOrErr(); err == nil {
		result.Image = fileUUIDToURLPath(i.FileID)
	}
	return result
}
func dbordersTopborders(where []string, od []*ent.Order) []*pb.Order {
	res := make([]*pb.Order, len(od))
	for i, o := range od {
		res[i] = dborderTopborder(where, o)
	}
	return res
}
func dborderTopborder(where []string, od *ent.Order) *pb.Order {
	tmp := make([]*pb.OrderField, len(od.Edges.Items))
	for i, item := range od.Edges.Items {
		tmp[i] = &pb.OrderField{
			Menu:   int64(item.Edges.Menu.ID),
			Count:  int32(item.Count),
			Amount: item.Edges.Menu.Price.Amount(),
			Name:   item.Edges.Menu.Name,
		}
	}
	return &pb.Order{
		Id:         int64(od.ID),
		Where:      pathTopbrestpath(where, ""),
		Who:        strconv.FormatInt(int64(od.Edges.Who.ID), 10),
		OrderAt:    od.OrderAt.Format(time.RFC3339),
		CookingAt:  timeOrEmpty(od.CookingAt),
		DeliverAt:  timeOrEmpty(od.DeliverAt),
		CompleteAt: timeOrEmpty(od.CompleteAt),
		State:      stateGen(od.CookingAt, od.DeliverAt, od.CompleteAt),
		Purchased:  tmp,
	}
}
func dbcategoriesTopbcategories(categories []*ent.Category) []string {

	opts := make([]string, len(categories))
	for i, cat := range categories {
		opts[i] = cat.Name
	}
	return opts
}
func dbrestTopbrest(r *ent.Restaurant) *pb.Restaurant {
	result := &pb.Restaurant{
		Name:                 pathTopbrestpath(getRecursiveName(r), r.URI),
		Description:          r.Description,
	}
	if avatar, err := r.Edges.AvatarOrErr(); err == nil {
		result.Image = fileUUIDToURLPath(avatar.FileID)
	}
	return result
}
func pathTopbrestpath(p []string, url string) *pb.RestaurantName {
	result := &pb.RestaurantName{
		Path:                 p,
		Url:                  url,
	}
	//if len(p) >=1 {
	//	result.Path1 = p[0]
	//}
	//if len(p) >=2 {
	//	result.Path2 = p[1]
	//}
	return result
}

func timeUnixOrRFC3339(t string) (at time.Time, empty bool, valid bool) {
	if t, err := time.Parse(time.RFC3339, t); err == nil {
		return t, false, true
	}
	if unixEpoch, err := strconv.ParseInt(t, 10, 64); err == nil {
		return time.Unix(unixEpoch, 0), false, true
	}
	if len(t) == 0 {
		return time.Unix(0, 0), true, true
	}
	return time.Unix(0, 0), false, true
}
func timeOrEmpty(t *time.Time) string {
	if t != nil {
		return t.Format(time.RFC3339)
	}
	return ""
}
func stateGen(cooking, deliver, complete *time.Time) pb.OrderState {
	if complete != nil {
		return pb.OrderState_Complete
	}
	if deliver != nil {
		return pb.OrderState_Deliver
	}
	if cooking != nil {
		return pb.OrderState_Cooking
	}
	return pb.OrderState_Pending
}

func randomStateString() string {
	var buff = make([]byte, 20)
	// state 문자열 생성
	_, err := rand.Read(buff)
	if err != nil {
		panic(errors.WithMessage(err, "randomStateString"))
	}
	return base64.URLEncoding.EncodeToString(buff)
}

func loadSession(client *redis.Client, ctx context.Context, state string) (session *Session) {
	session = new(Session)
	err := client.Get(ctx, state).Scan(session)
	if err != nil {
		return nil
	}
	return session
}
func saveSession(client *redis.Client, ctx context.Context, state string, session *Session, timeout time.Duration) (err error) {
	_, err = client.Set(ctx, state, session, timeout).Result()
	return err
}

func getRecursiveName(rest *ent.Restaurant) []string {
	if p, err := rest.Edges.ParentOrErr(); err == nil {
		return append(getRecursiveName(p), rest.Name)
	}
	return []string{rest.Name}
}
func noStore(c *gin.Context) {
	c.Header("Cache-Control", "no-store")
}
