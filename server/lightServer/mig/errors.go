package mig

import "github.com/pkg/errors"

var (
	ErrOpenSSL      = errors.New("openSSL Error")
	ErrNoOpenSSL    = errors.WithMessage(ErrOpenSSL, "No openssl")
	ErrRSAGenerate  = errors.New("X509 generation fail")
	ErrX509Generate = errors.New("X509 generation fail")
)
var (
	ErrInitialize = errors.New("Initialize fail")
)
var (
	ErrDB                   = errors.WithMessage(ErrInitialize, "DB Fail")
	ErrDBSource             = errors.WithMessage(ErrDB, "Source fail")
	ErrDBConnectedCloseFail = errors.WithMessage(ErrDB, "DB Opened connection closing fail")
	ErrDBConnectFail        = errors.WithMessage(ErrDB, "DB connect failure")
	ErrDBMigration          = errors.WithMessage(ErrDB, "Migration Fail")
)
var (
	ErrHost = errors.WithMessage(ErrInitialize, "Host Fail")
)

var (
	ErrX509Load   = errors.WithMessage(ErrInitialize, "X509 load fail")
	ErrX509Loaded = errors.WithMessage(ErrInitialize, "X509 loaded")
)

var (
	ErrOAuth2      = errors.WithMessage(ErrInitialize, "OAuth2")
	ErrOAuth2Exist = errors.WithMessage(ErrOAuth2, "OAuth2 Service, Exist add")
)

var (
	ErrEthereum      = errors.WithMessage(ErrInitialize, "Ethereum")
	ErrEthereumSetuped = errors.WithMessage(ErrEthereum, "ethereum setuped")
)
