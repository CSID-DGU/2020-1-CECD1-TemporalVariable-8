package mig

import (
	"bytes"
	"context"
	"encoding/gob"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-contrib/cache/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/hashicorp/go-multierror"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"lightServer/ent"
	"lightServer/ent/authorize"
	"lightServer/mig/migc"
	"math/big"
	"net"
	"net/http"
	"net/url"
	"time"
)

type (
	MakItGo struct {
		Configs Configs
		Unsafe  *unsafeServer
	}
	Configs struct {
		// Dynamics
		Name                  string
		DefaultDBTimeout      time.Duration
		DefaultSessionTimeout time.Duration
		// Static
		oauth2 map[string]*oauth2Config
		guest *ent.User
		url *url.URL
	}
	oauth2Config struct {
		conf       *oauth2.Config
		provider   authorize.Provider
		getProfile func(client *http.Client) (gjson.Result, error)
	}
	unsafeServer struct {
		DB              *ent.Client       // ent로 연결된 데이터베이스
		Engine          *gin.Engine       // HTTP 서버 라우팅
		HTTPServer      *http.Server      // HTTP 서버
		Net             net.Listener      // 네트워크 연결
		Logger          *zap.Logger       // 로깅용 객체
		FS              http.FileSystem   // 파일시스템 주소
		Ethereum        *ethclient.Client // 이더리움의 스마트 컨트랙트
		Contract        *migc.Migc        // 이더리움의 스마트 컨트랙트
		ContractAddress common.Address    // 이더리움의 스마트 컨트랙트
		Cache           *redis.Client     // 세션관리, 페이지 렌더링, 등등
	}
	binaryWrapper struct {
		i interface{}
	}
)

func (b *binaryWrapper) UnmarshalBinary(data []byte) error {
	return gob.NewDecoder(bytes.NewBuffer(data)).Decode(b.i)
}

func (b *binaryWrapper) MarshalBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(nil)
	err = gob.NewEncoder(buf).Encode(b.i)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s *unsafeServer) Get(key string, value interface{}) error {
	v, err := s.Cache.Get(context.Background(), key).Result()
	if err != nil {
		return err
	}
	return utils.Deserialize([]byte(v), value)
}

func (s *unsafeServer) Set(key string, value interface{}, expire time.Duration) error {

	return s.Cache.Set(context.Background(), key, &binaryWrapper{i: value}, expire).Err()
}

func (s *unsafeServer) Add(key string, value interface{}, expire time.Duration) error {
	return s.Cache.SetNX(context.Background(), key, &binaryWrapper{i: value}, expire).Err()
}

func (s *unsafeServer) Replace(key string, data interface{}, expire time.Duration) error {
	return s.Cache.SetXX(context.Background(), key, &binaryWrapper{i: data}, expire).Err()
}

func (s *unsafeServer) Delete(key string) error {
	return s.Cache.Del(context.Background(), key).Err()
}

func (s *unsafeServer) Increment(key string, data uint64) (uint64, error) {
	v, err := s.Cache.IncrBy(context.Background(), key, int64(data)).Result()
	return uint64(v), err

}

func (s *unsafeServer) Decrement(key string, data uint64) (uint64, error) {
	v, err := s.Cache.DecrBy(context.Background(), key, int64(data)).Result()
	return uint64(v), err
}

func (s *unsafeServer) Flush() error {
	return s.Cache.FlushAll(context.Background()).Err()
}

func (s *MakItGo) Run() error {
	defer s.Unsafe.Logger.Sync()
	s.Unsafe.Logger.Named("[ MakItGo ]").Info("MakItGo.Run")

	return s.Unsafe.HTTPServer.Serve(s.Unsafe.Net)
}

func (s *MakItGo) Sync() error{
	count, err := s.Unsafe.Contract.ListURLCount(nil)
	if err != nil {
		return err
	}
	s.Unsafe.Logger.Info("[ MakItGo ]", zap.String("message", "맛잇고 동기화 시작"))
	for i := 0; i < int(count); i++ {
		u, err := s.Unsafe.Contract.Urls(nil, big.NewInt(int64(i)))
		s.Unsafe.Logger.Warn("[ MakItGo ]",
			zap.String("message", "URL 감지"),
			zap.String("url", u))
		if err != nil {
			return err
		}

	}
	s.Unsafe.Logger.Info("[ MakItGo ]", zap.String("message", "맛잇고 동기화 완료"))

	return nil
}
func (s *MakItGo) DebugString() {
	s.Unsafe.Logger.Info("[ MakItGo ]", zap.String("message", "맛잇고 동기화 시작"))
	s.Unsafe.Logger.Warn("[ MakItGo ]", zap.String("message", "'192.168.0.8' 감지"))
	//time.Sleep(400 * time.Millisecond)
	//s.Unsafe.Logger.Warn("[ MakItGo ]", zap.String("message", "'/동국반점' 감지"))
	//time.Sleep(100 * time.Millisecond)
	//s.Unsafe.Logger.Warn("[ MakItGo ]", zap.String("message", "'/동국반점' 동기화"), zap.Int("메뉴 동기화", 3))
	//s.Unsafe.Logger.Warn("[ MakItGo ]", zap.String("message", "'/동국반점/충무로점' 감지"))
	//time.Sleep(120 * time.Millisecond)
	//s.Unsafe.Logger.Warn("[ MakItGo ]", zap.String("message", "'/동국반점/충무로점' 동기화"), zap.Int("메뉴 동기화", 4))
	time.Sleep(600 * time.Millisecond)
	s.Unsafe.Logger.Info("[ MakItGo ]", zap.String("message", "맛잇고 동기화 완료"))
}
func (s *MakItGo) DebugString2() {
	s.Unsafe.Logger.Info("[ MakItGo ]", zap.String("message", "맛잇고 동기화 시작"))
	s.Unsafe.Logger.Warn("[ MakItGo ]", zap.String("message", "'/동국반점' 공개됨"))
	s.Unsafe.Logger.Warn("[ MakItGo ]", zap.String("message", "'/동국반점/충무로점' 공개됨"))
	s.Unsafe.Logger.Info("[ MakItGo ]", zap.String("message", "맛잇고 동기화 완료"))
}


func (s *MakItGo) RunChannel(c chan error) <-chan error {
	go func() {
		err := s.Run()
		if err != nil {
			c <- err
		}
		close(c)
	}()
	return c
}

func (s *MakItGo) Close() (err error) {
	if dbErr := s.Unsafe.DB.Close(); dbErr != nil {
		multierror.Append(err, dbErr)
	}
	if cacheErr := s.Unsafe.Cache.Close(); cacheErr != nil {
		multierror.Append(err, cacheErr)
	}
	if httpErr := s.Unsafe.HTTPServer.Close(); httpErr != nil {
		multierror.Append(err, httpErr)
	}
	s.Unsafe.Ethereum.Close()
	return err
}
