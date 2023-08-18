package scheduler

import (
	"area/config"
	"area/interfaces/data"
	"common/entity/area"
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"os"
	"time"
)

type AreaScheduler struct {
	path        string
	db          *gorm.DB
	logger      *logrus.Logger
	processTime time.Time
}

func (a *AreaScheduler) Process() error {
	var result data.HangJeongDong
	//데이터 읽어와서 인서트 처리
	// 읽어오는거 테스트
	file, err := os.ReadFile(a.path + "/HangJeongDong_ver20230701.json")
	if err != nil {
		a.logger.WithError(err).Error("config file read failure")
		return err
	}
	err = json.Unmarshal(file, &result)
	if err != nil {
		a.logger.WithError(err).Error("json unmarshal failure")
		return err
	}
	//var add []area.SiggArea

	for _, v := range result.Features {
		if err = a.db.Transaction(func(tx *gorm.DB) error {

			loc := &area.SiggArea{
				AreaId:   v.Properties.Sgg,
				CityId:   v.Properties.Sido,
				AmdCd:    v.Properties.AdmCd,
				SiggName: v.Properties.Sggnm,
				Sido: area.SidoArea{
					CityId:   v.Properties.Sido,
					SidoName: v.Properties.SidoName,
				},
			}
			return a.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(loc).Error
		}); err != nil {
			a.logger.WithError(err).Error("transaction failure")
			break
		}
	}

	return nil
}

func (a *AreaScheduler) Run(ctx context.Context) <-chan error {
	batch := make(chan error)
	go func() {
		defer close(batch)
		defer fmt.Println("area scheduler end")
		ticker := time.NewTicker(24 * time.Hour)
		//ticker := time.Tick(5 * time.Second)
		for {
			select {
			case <-ctx.Done():
				batch <- ctx.Err()
				return
			case <-ticker.C:
				a.logger.Info("area scheduler running")
				if err := a.Process(); err != nil {
					batch <- err
				}
			}
		}
	}()
	return batch
}

func NewAreaScheduler(db *gorm.DB, cfg *config.AreaConfig, log *logrus.Logger) *AreaScheduler {
	return &AreaScheduler{
		path:   cfg.Scheduler.Path,
		db:     db,
		logger: log,
	}
}
