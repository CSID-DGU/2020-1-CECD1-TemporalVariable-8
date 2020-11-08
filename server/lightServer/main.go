package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis/v8"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"lightServer/mig"
	"os"
	"os/signal"
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
	sigin := make(chan os.Signal)
	signal.Notify(sigin, os.Interrupt, os.Kill)
	//
	cfg := config()
	//
	logfile, err := os.OpenFile(fmt.Sprintf("./MakItGo/%s.log", time.Now().Format("2006-01-02")), os.O_APPEND|os.O_CREATE|os.O_SYNC|os.O_WRONLY, os.FileMode(0644))
	if err != nil {
		panic(err)
	}
	defer logfile.Close()
	//
	server, err := mig.BuildServer().
		Host("").
		LogConsole().
		LogFile(logfile).
		Redis(&redis.Options{
			Addr: "localhost:6379",
			DB:   0,
		}).
		Google(
			cfg.Get("oauth2.google.id").String(),
			cfg.Get("oauth2.google.secret").String(),
		).
		DialSmartContract("http://localhost:8545", common.HexToAddress("0x9e7495ba2dd6Cc064D6315511cCDA0aF32d3F966")).
		DBDriver(mig.SQLITE3, "file:./MakItGo/mig.db").
		Build()

	if err != nil {
		panic(err)
	}
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
}
