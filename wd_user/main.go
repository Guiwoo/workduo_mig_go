package main

import (
	"flag"
	"fmt"
	"wd_common/wd_database"
	"wd_user/pkg/config"
)

var (
	//실행하는 파일 의 go.mod 기준으로 경로 탐색
	configFile = flag.String("cfg", "./pkg/config/dev.yaml", "The path to the config file")
	cfg        config.User
)

func main() {
	flag.Parse()
	cfg.ParseConfig(*configFile)
	cfg.PrintConfig()

	db, err := wd_database.Open(cfg.Dsn)
	if err != nil {
		panic(err)
	}

	fmt.Println(db)
}
