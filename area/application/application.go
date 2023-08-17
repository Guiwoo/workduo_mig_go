package application

import (
	"area/config"
	"area/interfaces/restapis"
	"common/database"
	"common/pipeline"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
)

type AreaService struct {
	Config    *config.AreaConfig
	Logger    *logrus.Logger
	AreaDB    *gorm.DB
	Pool      database.DBPool
	APIServer *restapis.APIService
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

func (a *AreaService) SetUp(cfg *config.AreaConfig) {
	a.setUpDB(cfg)
	a.APIServer = restapis.NewAPIService(cfg.Server.Port)
}

func (a *AreaService) Run() error {
	stream := pipeline.NewStream()
	stream.Add(a.APIServer.Run)
	return stream.Run()
}

func NewAreaService(cfg *config.AreaConfig) *AreaService {
	service := &AreaService{
		Config: cfg,
	}
	service.SetUp(cfg)
	return service
}
