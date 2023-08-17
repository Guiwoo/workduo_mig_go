package application

import (
	"area/config"
	"common/database"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
)

type AreaService struct {
	Config *config.AreaConfig
	Logger *logrus.Logger
	AreaDB *gorm.DB
	Pool   database.DBPool
}

func (a *AreaService) setUpDB(cfg *config.AreaConfig) {
	a.Pool = database.NewDBPool()

	db, err := a.Pool.ConnectArea(cfg.Database.Databases[database.DBAREA], cfg.Database.LogLevel, cfg.Database.MaxIdle, cfg.Database.MaxConn, cfg.Database.MaxLifeCycle)
	if err != nil {
		logrus.WithError(err).Error("database connect failure")
		os.Exit(1)
	}
	a.AreaDB = db
}

func (a *AreaService) Run() error {
	return nil
}

func NewAreaService(cfg *config.AreaConfig) *AreaService {
	service := &AreaService{
		Config: cfg,
	}
	service.setUpDB(cfg)
	return service
}
