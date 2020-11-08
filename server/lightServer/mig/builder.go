package mig

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/biter777/countries"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io"
	"io/ioutil"
	"lightServer/ent"
	"lightServer/mig/migc"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"time"
)

type serverBuilder struct {
	//
	cert   tls.Certificate
	zcores []zapcore.Core
	//
	svr *MakItGo
	err error
}

func BuildServer() *serverBuilder {
	eg := gin.New()
	mig := &MakItGo{
		Unsafe: &unsafeServer{
			Engine: eg,
			DB:     nil,
			HTTPServer: &http.Server{
				Addr:              "",
				Handler:           eg,
				TLSConfig:         nil,
				ReadTimeout:       0,
				ReadHeaderTimeout: 0,
				WriteTimeout:      0,
				IdleTimeout:       0,
				MaxHeaderBytes:    0,
				TLSNextProto:      nil,
				ConnState:         nil,
				ErrorLog:          nil,
			},
			Cache: nil,
			FS:    http.Dir("."),
		},
		Configs: Configs{
			Name:                  "",
			DefaultDBTimeout:      90 * time.Second,
			DefaultSessionTimeout: 60 * time.Minute,
			oauth2:                make(map[string]*oauth2.Config),
		},
	}
	mig.Unsafe.HTTPServer.BaseContext = func(listener net.Listener) context.Context {
		return context.WithValue(context.Background(), MAK_IT_GO, mig)
	}
	bd := &serverBuilder{
		svr: mig,
		err: nil,
	}
	bd.H2C()
	return bd
}
func (s *serverBuilder) Default() *serverBuilder {
	return s.
		Host("").
		DBDriver(SQLITE3, "file::memory:?mode=memory&cache=shared&_fk=1").
		Redis(&redis.Options{
			Addr: "localhost:6379",
			DB:   0,
		}).
		H2C()
}
func (s *serverBuilder) LogConsole() *serverBuilder {
	if s.err != nil {
		return s
	}
	consoleEnc := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	s.zcores = append(s.zcores,
		zapcore.NewCore(
			consoleEnc,
			zapcore.Lock(os.Stdout),
			zap.LevelEnablerFunc(func(level zapcore.Level) bool {
				return level >= zapcore.InfoLevel
			}),
		),
	)
	return s
}
func (s *serverBuilder) LogFile(file io.Writer) *serverBuilder {
	if s.err != nil {
		return s
	}
	jsonEnc := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	s.zcores = append(s.zcores,
		zapcore.NewCore(
			jsonEnc,
			zapcore.AddSync(file),
			zap.LevelEnablerFunc(func(level zapcore.Level) bool { return level >= zapcore.ErrorLevel }),
		),
	)
	return s
}

func (s *serverBuilder) Host(addr string) *serverBuilder {
	if s.err != nil {
		return s
	}
	if len(addr) <= 0 {
		addr = ":http"
	}
	_, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		s.err = errors.WithMessage(ErrHost, err.Error())
		return s
	}
	s.svr.Unsafe.HTTPServer.Addr = addr
	return s
}
func (s *serverBuilder) TLSFile(key, pem string) *serverBuilder {
	if s.err != nil {
		return s
	}
	var cert tls.Certificate
	cert, s.err = tls.LoadX509KeyPair(key, pem)
	if s.err != nil {
		s.err = errors.WithMessage(ErrX509Load, s.err.Error())
		return s
	}
	return s.TLS(cert)
}
func (s *serverBuilder) TLS(certificate tls.Certificate) *serverBuilder {
	if s.err != nil {
		return s
	}
	if s.svr.Unsafe.HTTPServer.TLSConfig == nil {
		s.svr.Unsafe.HTTPServer.TLSConfig = new(tls.Config)
	}
	if _, ok := s.svr.Unsafe.HTTPServer.Handler.(*gin.Engine); !ok {
		s.svr.Unsafe.HTTPServer.Handler = s.svr.Unsafe.Engine
	}
	if len(s.svr.Unsafe.HTTPServer.TLSConfig.Certificates) > 0 {
		s.err = ErrX509Loaded
		return s
	}
	// https, http2 기능 켜기 위함
	s.svr.Unsafe.HTTPServer.TLSConfig.NextProtos = append(s.svr.Unsafe.HTTPServer.TLSConfig.NextProtos, "http/1.1", "h2")
	s.svr.Unsafe.HTTPServer.TLSConfig.Certificates = []tls.Certificate{certificate}
	return s
}
func (s *serverBuilder) H2C() *serverBuilder {
	if s.err != nil {
		return s
	}
	if s.svr.Unsafe.HTTPServer.TLSConfig != nil {
		s.svr.Unsafe.HTTPServer.TLSConfig = nil
	}
	if _, ok := s.svr.Unsafe.HTTPServer.Handler.(*gin.Engine); ok {
		s.svr.Unsafe.HTTPServer.Handler = h2c.NewHandler(s.svr.Unsafe.Engine, &http2.Server{})
	}
	return s
}
func (s *serverBuilder) DBDriver(driver DBDriver, source string) *serverBuilder {
	if s.err != nil {
		return s
	}
	if s.svr.Unsafe.DB != nil {
		err := s.svr.Unsafe.DB.Close()
		if err != nil {
			s.err = errors.WithMessage(ErrDBConnectedCloseFail, err.Error())
			return s
		}
	}
	if driver == SQLITE3 {
		parsed, err := url.Parse(source)
		if err != nil {
			s.err = errors.WithMessage(ErrDBSource, err.Error())
			return s
		}
		parsed.Scheme = "file"
		q := parsed.Query()
		q.Set("_fk", "1")
		parsed.RawQuery = q.Encode()
		source = parsed.String()
	}
	s.svr.Unsafe.DB, s.err = ent.Open(string(driver), source)
	if s.err != nil {
		s.err = errors.WithMessagef(ErrDBConnectFail, "%e : %s", s.err, source)
		return s
	}
	s.err = s.svr.Unsafe.DB.Schema.Create(context.Background())
	if s.err != nil {
		s.err = errors.WithMessage(ErrDBMigration, s.err.Error())
		return s
	}
	return s
}
func (s *serverBuilder) Redis(options *redis.Options) *serverBuilder {
	if s.err != nil {
		return s
	}
	s.svr.Unsafe.Cache = redis.NewClient(options)
	tout, _ := context.WithTimeout(context.Background(), 2*time.Second)
	s.err = s.svr.Unsafe.Cache.Ping(tout).Err()
	if s.err != nil {
		s.svr.Unsafe.Cache = nil
		return s
	}
	return s
}
func (s *serverBuilder) DialSmartContract(backendDial string, addr common.Address) *serverBuilder {
	if s.err != nil {
		return s
	}
	s.svr.Unsafe.Ethereum, s.err = ethclient.Dial(backendDial)
	if s.err != nil {
		return s
	}
	if s.svr.Unsafe.Contract != nil {
		s.err = ErrEthereumSetuped
		return s
	}
	s.svr.Unsafe.Contract, s.err = migc.NewMigc(addr, s.svr.Unsafe.Ethereum)
	if s.err != nil {
		return s
	}
	s.svr.Unsafe.ContractAddress = addr
	return s
}
//
func (s *serverBuilder) OAuth2(serviceName string, endpoint oauth2.Endpoint, clientID, ClientSecret string, scope ...string) *serverBuilder {
	if _, ok := s.svr.Configs.oauth2[serviceName]; ok {
		s.err = ErrOAuth2Exist
		return s
	}
	s.svr.Configs.oauth2[serviceName] = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: ClientSecret,
		Endpoint:     endpoint,
		Scopes:       scope,
	}
	return s
}

func (s *serverBuilder) Google(clientID, ClientSecret string) *serverBuilder {
	s.OAuth2("google", google.Endpoint, clientID, ClientSecret,
		"openid",
		"profile",
		"email",
	)
	return s
}

//
func (s *serverBuilder) Build() (*MakItGo, error) {
	if s.err != nil {
		return nil, s.err
	}
	var (
		isHTTP2 = false
	)
	//
	var err error
	// Gin 서버
	err = s.svr.buildGin()
	if err != nil {
		s.err = err
		return nil, s.err
	}

	//
	tcpaddr, err := net.ResolveTCPAddr("tcp", s.svr.Unsafe.HTTPServer.Addr)
	if err != nil {
		return nil, s.err
	}
	s.svr.Unsafe.Net, err = net.ListenTCP("tcp", tcpaddr)
	if err != nil {
		s.err = err
		return nil, s.err
	}
	if s.svr.Unsafe.HTTPServer.TLSConfig != nil {
		isHTTP2 = true
		s.svr.Unsafe.Net = tls.NewListener(s.svr.Unsafe.Net, s.svr.Unsafe.HTTPServer.TLSConfig)
	}
	// oauth2
	for service, config := range s.svr.Configs.oauth2 {
		redirect := url.URL{
			Scheme:  "http",
			Host:    s.svr.Unsafe.HTTPServer.Addr,
			RawPath: fmt.Sprintf("/auth/callback/%s", service),
		}
		if isHTTP2 {
			redirect.Scheme = "https"
		}
		config.RedirectURL = redirect.String()
	}
	// 로거
	s.svr.Unsafe.Logger = zap.New(zapcore.NewTee(s.zcores...))
	return s.svr, nil
}

type (
	openSSL      string
	X509Subjects struct {
		Country              countries.CountryCode // C
		StateName            string                // ST
		LocalName            string                // L
		OrganizationName     string                // O
		OrganizationUnitName string                // OU
		CommonName           string                // CN
	}
)

var OpenSSL openSSL = ""

func (o *X509Subjects) String() string {
	if o == nil {
		return "/"
	}
	return fmt.Sprintf("/C=%s/ST=%s/L=%s/O=%s/OU=%s/CN=%s", o.Country.Alpha2(), o.StateName, o.LocalName, o.OrganizationName, o.OrganizationUnitName, o.CommonName)
}
func (o openSSL) String() string {
	opensslpath := "openssl"
	if len(o) != 0 {
		opensslpath = string(o)
	}
	return opensslpath
}
func (o openSSL) Available() bool {
	_, err := exec.Command(o.String(), "version").Output()
	return err == nil
}
func (o openSSL) GenRSABytes(pass string) ([]byte, error) {
	if !o.Available() {
		return nil, ErrNoOpenSSL
	}
	keytxt, keyerr := exec.Command(
		o.String(),
		"genrsa", "-aes256",
		// 비밀번호
		"-passout", "pass:"+pass,
		// 2048 비트 키
		"2048").Output()
	if keyerr != nil {
		return nil, errors.WithMessage(ErrRSAGenerate, keyerr.Error())
	}
	return keytxt, nil
}
func (o openSSL) NewX509Bytes(pass string, rsa []byte, subject *X509Subjects) (x509 []byte, err error) {
	if !o.Available() {
		return nil, ErrNoOpenSSL
	}
	f, err := ioutil.TempFile("", "*.key")
	if err != nil {
		return nil, err
	}
	defer os.Remove(f.Name())
	defer f.Close()
	_, err = f.Write(rsa)
	if err != nil {
		return nil, err
	}
	pemtxt, pemerr := exec.Command(
		"openssl",
		"req", "-new", "-x509", "-sha256",
		// 비밀번호
		"-passin", "pass:"+pass,
		// 서브젝트
		"-subj", subject.String(),
		// 키 파일
		"-key", f.Name(),
		//10년
		"-days", "3650").Output()
	if pemerr != nil {
		return nil, errors.WithMessage(ErrX509Generate, pemerr.Error())
	}
	return pemtxt, nil
}
