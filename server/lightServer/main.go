package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"lightServer/mig"
	"os"
	"os/signal"
	"path"
	"time"
)

func config() gjson.Result {
	f, err := os.Open("./MakItGo/config.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fstr, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return gjson.ParseBytes(fstr)
}
func main() {
	gin.SetMode(gin.DebugMode)
	sigin := make(chan os.Signal)
	signal.Notify(sigin, os.Interrupt, os.Kill)
	//
	cfg := config()
	//

	logfile, err := os.OpenFile(path.Join(cfg.Get("logging.path").String(), fmt.Sprintf("./%s.log", time.Now().Format("2006-01-02"))), os.O_APPEND|os.O_CREATE|os.O_SYNC|os.O_WRONLY, os.FileMode(0644))
	if err != nil {
		panic(err)
	}
	defer logfile.Close()
	//
	//
	server, err := mig.BuildServer().
		Host("").
		LogConsole().
		LogFile(logfile).
		Redis(&redis.Options{
			Addr: cfg.Get("cache.address").String(),
			DB:   0,
		}).
		DNS(cfg.Get("dns").String()).
		Google(
			cfg.Get("oauth2.google.id").String(),
			cfg.Get("oauth2.google.secret").String(),
		).
		DialSmartContract(cfg.Get("blockchain.api_server").String(), common.HexToAddress(cfg.Get("blockchain.smart_contract").String())).
		DBDriver(mig.DBDriver(cfg.Get("db.driver").String()), cfg.Get("db.source").String()).
		Build()
	if err != nil {
		panic(err)
	}
	//err = server.Sync()
	//if err != nil {
	//	panic(err)
	//}
	server.DebugString()
LOOP:
	for {
		select {
		case err := <-server.RunChannel(make(chan error)):
			if err != nil {
				// TODO : Error Logging
			}
			break LOOP
		case sig := <-sigin:
			switch sig {
			case os.Interrupt:
				fmt.Println("Interrupt")
				err := server.Close()
				if err != nil {
					panic(err)
				}
			case os.Kill:
				fmt.Println("Kill")
				err := server.Close()
				if err != nil {
					panic(err)
				}
			}
		}
	}
	for i := 0; i < 100;i++ {
		fmt.Println()
	}
}
