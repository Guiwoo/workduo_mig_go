package main

import (
	"area/application"
	"area/config"
	"flag"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	//실행하는 파일 의 go.mod 기준으로 경로 탐색
	configFile = flag.String("cfg", "./config/config.yaml", "The path to the config file.")
)

func main() {
	flag.Parse()

	cfg := config.AreaConfig{}
	if err := cfg.LoadConfig(*configFile); err != nil {
		logrus.WithError(err).Error("config load failure")
		os.Exit(0)
	}

	app := application.NewAreaService(&cfg)

	if err := app.Run(); err != nil {
		logrus.WithError(err).Error("application run failure")
		os.Exit(1)
	}
}
