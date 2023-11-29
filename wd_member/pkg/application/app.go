package application

import (
	"gorm.io/gorm"
	"wd_common/wd_log"
	"wd_common/wd_streams"
	"wd_user/pkg/config"
)

type App struct {
	Server *UserRestAPI
}

func (app *App) Run() error {
	stream := wd_streams.New()
	stream.Add(app.Server.Run)
	return stream.Run()
}

func New(db *gorm.DB, cfg config.User) *App {
	log := wd_log.New(cfg.Log.Path, cfg.Log.FileName, cfg.Log.Level)
	return &App{
		Server: NewUserRestAPI(db, cfg.Listen, log),
	}
}
