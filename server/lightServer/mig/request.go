package mig

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gin-contrib/cache"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"lightServer/ent"
	"lightServer/ent/file"
	"lightServer/ent/predicate"
	"lightServer/ent/restaurant"
	"lightServer/mig/zaputil"
	"lightServer/pb"
	"math/big"
	"math/rand"
	"mime"
	"net/http"
	"path"
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
		}else{
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
	// /oauth2/self
	s.Unsafe.Engine.GET("/auth/self", func(c *gin.Context) {
		var (
			cState, err = c.Cookie("state")
		)
		if err != nil {
			c.JSON(http.StatusOK, nil)
			return
		}
		var (
			session = mustSession(s.Unsafe.Cache, context.Background(), cState, s.Configs.DefaultSessionTimeout)
		)
		c.JSON(http.StatusOK, session)
	})
	// auth
	// /auth/login/:target
	// 동작 : :target을 통해 OAuth2 로그인을 시도함
	s.Unsafe.Engine.GET("/auth/login/:target", func(c *gin.Context) {
		var (
			oauth2config *oauth2.Config
			ok           bool
			state        = randomStateString()
		)
		if oauth2config, ok = s.Configs.oauth2[c.Query("target")]; !ok {
			// TODO
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		//
		c.Redirect(http.StatusTemporaryRedirect, oauth2config.AuthCodeURL(state))
	})
	// auth
	// /auth/callback/:target
	// 동작 : :target을 통해 OAuth2 로그인을 시도한 후 그 결과를 받아 옴
	s.Unsafe.Engine.GET("/auth/callback/:target", func(c *gin.Context) {
		var (
			qState = c.Query("state")
			qCode  = c.Query("code")
		)
		var (
			oauth2config *oauth2.Config
			ok           bool
			session      = mustSession(s.Unsafe.Cache, context.Background(), qState, s.Configs.DefaultSessionTimeout)
		)
		if oauth2config, ok = s.Configs.oauth2[c.Query("target")]; !ok {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		tk, err := oauth2config.Exchange(context.Background(), qCode)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		//
		session.Token = tk
		err = saveSession(s.Unsafe.Cache, context.Background(), qState, session, s.Configs.DefaultSessionTimeout)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		//
		c.SetCookie("state", qState, int(s.Configs.DefaultSessionTimeout.Seconds()), "/", "", false, true)
		c.Status(http.StatusOK)
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
	s.Unsafe.Engine.Group("/file").
		GET("/:fid", func(c *gin.Context) {
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
	restful := s.Unsafe.Engine.Group("/api")
	// RESTful API : restaurant
	// /api/restaurants/
	// 동작 : 레스토랑 목록에 대해 전달
	restful.GET("/restaurants", cache.CachePage(s.Unsafe, 1*time.Minute, func(ctx *gin.Context) {
		timedCtx, _ := context.WithTimeout(context.Background(), s.Configs.DefaultDBTimeout)
		names, err := s.Unsafe.DB.Restaurant.Query().Select(restaurant.FieldName).Strings(timedCtx)
		if err != nil {
			panic(err)
		}
		renderByTyped(ctx, http.StatusOK, &pb.RestaurantStatistics{
			Names: names,
		})
	}))
	// RESTful API : restaurant
	// 라우터
	restfulRSub := restful.Group("/restaurants/:name/:subname", func(c *gin.Context) {
		var (
			qName    = c.Param("name")
			qSubname = c.Param("subname")
		)
		var (
			name    = restaurant.Name(qName)
			subName predicate.Restaurant
		)

		fmt.Println(qName, qSubname)
		switch qSubname {
		case "":
			subName = restaurant.SubNameIsNil()
		default:
			subName = restaurant.SubName(qSubname)
		}
		r, err := s.Unsafe.DB.Restaurant.Query().
			WithRoot().WithParent().WithChildren().
			Where(name, subName).Only(context.Background())
		if err != nil {
			// TODO Error
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.Set(restaurant.Table, r)
	})
	// RESTful API : restaurant
	// /api/restaurants/:name/:subname/?type=타입
	// 동작 : 특정 :name, :subname 지점의 정보를 가져옴
	// 타입 : json | protobuf | default=json, 출력하는 데이터의 포멧 형식
	restfulRSub.GET("/", cache.CachePage(s.Unsafe, 1*time.Minute, func(c *gin.Context) {
		tb := c.MustGet(restaurant.Table).(*ent.Restaurant)
		pbobj := &pb.Restaurant{
			Name: tb.Name,
		}
		if len(tb.SubName) > 0 {
			pbobj.SubName = tb.SubName
		}
		if tb.URI != nil {
			pbobj.Url = tb.URI.String()
		}
		if tb.Edges.Root != nil {
			pbobj.Root = tb.Edges.Root.Name
		}
		if tb.Edges.Parent != nil {
			pbobj.Parent = tb.Edges.Parent.Name
		}
		pbobj.Children = make([]string, 0, len(tb.Edges.Children))
		for _, c := range tb.Edges.Children {
			pbobj.Children = append(pbobj.Children, c.Name)
		}
		renderByTyped(c, http.StatusOK, pbobj)
	}))
	// RESTful API : restaurant
	// /api/restaurants/:name/:subname/menus
	// 동작 : 특정 :name, :subname 지점의 메뉴를 모두 가져옴
	restfulRSub.GET("/menus", cache.CachePage(s.Unsafe, 10*time.Minute, func(c *gin.Context) {
		tb := c.MustGet(restaurant.Table).(*ent.Restaurant)
		menus, err := tb.QueryMenus().WithOwner().WithOptions().All(context.Background())
		if err != nil {
			s.Unsafe.Logger.Error("[ DB ]",
				zap.String("message", "Menus Error"),
				zap.String("error", err.Error()),
			)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		//
		tmp := &pb.Menus{
			Owner: tb.Name,
			Menus: dbmenusTopbmenus(menus),
		}
		renderByTyped(c, http.StatusOK, tmp)
	}))
	// Not implemented
	// RESTful API : restaurant
	// /api/restaurants/:name/:subname/field/:field
	// 동작 : 특정 :name, :subname 지점의 :field를 읽음
	//restfulRSub.GET("/field/:field", func(c *gin.Context) {
	//	tb := c.MustGet(restaurant.Table).(*ent.Restaurant)
	//	switch c.Param("field") {
	//	case restaurant.FieldName:
	//		c.String(http.StatusOK, tb.Name)
	//	case restaurant.FieldID:
	//		c.String(http.StatusOK, strconv.Itoa(tb.ID))
	//	case restaurant.FieldURI:
	//		c.String(http.StatusOK, tb.URI.String())
	//	case restaurant.EdgeRoot:
	//		if tb.Edges.Root != nil {
	//			c.String(http.StatusOK, tb.Edges.Root.Name)
	//		} else {
	//			c.AbortWithStatus(http.StatusNotFound)
	//		}
	//	case restaurant.EdgeParent:
	//		if tb.Edges.Parent != nil {
	//			c.String(http.StatusOK, tb.Edges.Parent.Name)
	//		} else {
	//			c.AbortWithStatus(http.StatusNotFound)
	//		}
	//		c.String(http.StatusOK, tb.URI.String())
	//	case restaurant.EdgeChildren:
	//		if tb.Edges.Children != nil || len(tb.Edges.Children) > 0 {
	//			res := make([]string, len(tb.Edges.Children))
	//			for i, child := range tb.Edges.Children {
	//				res[i] = child.Name
	//			}
	//			c.JSON(http.StatusOK, res)
	//		} else {
	//			c.AbortWithStatus(http.StatusNotFound)
	//		}
	//		c.String(http.StatusOK, tb.URI.String())
	//	default:
	//		c.AbortWithStatus(http.StatusNotFound)
	//	}
	//
	//})

	//adminc := restful.Group("/admin")
	//adminc.GET("")
	return nil
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

func dbmenusTopbmenus(menus []*ent.Menu) []*pb.Menu {
	ret := make([]*pb.Menu, len(menus))
	for i, menu := range menus {
		ret[i] = dbmenuTopbmenu(menu)
	}
	return ret
}

func dbmenuTopbmenu(menu *ent.Menu) *pb.Menu {
	opts := make([]*pb.Menu, len(menu.Edges.Options))
	for i, mopt := range menu.Edges.Options {
		opts[i] = dbmenuTopbmenu(mopt)
	}
	return &pb.Menu{
		Name:        menu.Name,
		Description: menu.Description,
		PriceAmount: menu.Price.Amount(),
		Currency:    menu.Price.Currency().Code,
		Owner:       menu.Edges.Owner.Name,
		Options:     opts,
	}
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
func mustSession(client *redis.Client, ctx context.Context, state string, timeout time.Duration) (session *Session) {
	err := client.Get(ctx, state).Scan(session)
	if err != nil {
		session = new(Session)
		err = client.Set(ctx, state, session, timeout).Err()
		if err != nil {
			panic(err)
		}
	}
	return session
}
func saveSession(client *redis.Client, ctx context.Context, state string, session *Session, timeout time.Duration) (err error) {
	return client.SetXX(ctx, state, session, timeout).Err()
}
