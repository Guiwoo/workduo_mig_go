package application

import (
	"gorm.io/gorm"
	"wd_area/pkg/config"
	"wd_common/wd_database"
	"wd_common/wd_log"
	"wd_common/wd_streams"
)

type App struct {
	Server *AreaRestAPI
}

func (app *App) Run() error {
	stream := wd_streams.New()
	stream.Add(app.Server.Run)
	return stream.Run()
}

func New(db *gorm.DB, cfg config.Area) *App {
	db, err := wd_database.Open(cfg.Dsn)
	if err != nil {
		panic(err)
	}

	log := wd_log.New(cfg.Log.Path, cfg.Log.FileName, cfg.Log.Level)
	return &App{
		Server: NewAreaRestAPI(db, cfg.Listen, log),
	}
}
