package main

import (
	"area"
	"common/database"
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	//실행하는 파일 의 go.mod 기준으로 경로 탐색
	configFile = flag.String("cfg", "./cmd/config.yaml", "The path to the config file.")
)

func main() {
	flag.Parse()

	cfg := area.AreaConfig{}
	if err := cfg.LoadConfig(*configFile); err != nil {
		logrus.WithError(err).Error("config load failure")
		os.Exit(0)
	}

	pool := database.NewDBPool()

	db, err := pool.ConnectArea(cfg.Database.Databases[database.DBAREA], cfg.Database.LogLevel, cfg.Database.MaxIdle, cfg.Database.MaxConn, cfg.Database.MaxLifeCycle)
	if err != nil {
		logrus.WithError(err).Error("database connect failure")
		os.Exit(1)
	}

	fmt.Println(db)
}
