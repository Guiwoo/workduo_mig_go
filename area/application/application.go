package application

import (
	"area/application/scheduler"
	"area/application/service"
	"area/config"
	"area/interfaces/restapis"
	"common/database"
	"common/pipeline"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
)

type AreaService struct {
	Config          *config.AreaConfig
	Logger          *logrus.Logger
	AreaDB          *gorm.DB
	Pool            database.DBPool
	TemplateService *restapis.TemplateService
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
	a.TemplateService = restapis.NewAPIService(cfg.Server.Port, service.NewService(a.AreaDB))
}

func (a *AreaService) Run() error {
	areaScheduler := scheduler.NewAreaScheduler(a.AreaDB, a.Config, a.Logger)
	stream := pipeline.NewStream()
	stream.Add(a.TemplateService.Run)
	stream.Add(areaScheduler.Run)
	return stream.Run()
}

func NewAreaService(cfg *config.AreaConfig) *AreaService {
	svc := &AreaService{
		Config: cfg,
		Logger: logrus.StandardLogger(),
	}
	svc.SetUp(cfg)
	return svc
}
