package application

import (
	"context"
	"encoding/json"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"os"
	"time"
	"wd_area/pkg/model"
)

type Feature struct {
	Type       string `json:"type"`
	Properties struct {
		AdmNm    string `json:"adm_nm"`
		AdmCd    string `json:"adm_cd"`
		AdmCd2   string `json:"adm_cd2"`
		Sgg      string `json:"sgg"`
		Sido     string `json:"sido"`
		SidoName string `json:"sidonm"`
		Temp     string `json:"temp"`
		Sggnm    string `json:"sggnm"`
		AdmCd8   string `json:"adm_cd8"`
	} `json:"properties"`
}

type HangJeongDong struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

type Scheduler struct {
	db            *gorm.DB
	log           zerolog.Logger
	path          string
	processTime   time.Time
	nextExecution time.Time
	lastExecution time.Time
}

func (service *Scheduler) insertData() error {
	var result HangJeongDong
	//데이터 읽어와서 인서트 처리
	file, err := os.ReadFile(service.path + "/HangJeongDong_ver20230701.json")
	if err != nil {
		service.log.Error().Err(err).Msg("fail to read file")
		return err
	}
	err = json.Unmarshal(file, &result)
	if err != nil {
		service.log.Error().Err(err).Msg("json unmarshal failure")
		return err
	}
	//var add []area.SiggArea

	for _, v := range result.Features {
		if err = service.db.Transaction(func(tx *gorm.DB) error {

			loc := &model.SiggArea{
				AreaId:   v.Properties.Sgg,
				CityId:   v.Properties.Sido,
				AmdCd:    v.Properties.AdmCd,
				SiggName: v.Properties.Sggnm,
				Sido: model.SidoArea{
					CityId:   v.Properties.Sido,
					SidoName: v.Properties.SidoName,
				},
			}
			return service.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(loc).Error
		}); err != nil {
			service.log.Error().Err(err).Msg("transaction failure")
			break
		}
	}

	return nil
}

func (service *Scheduler) executionTime() {
	t := service.lastExecution
	loc, _ := time.LoadLocation("Asia/Seoul")
	dateTime := time.Date(t.Year(), t.Month(), t.Day(), 5, 0, 0, 0, loc)
	service.nextExecution = dateTime.AddDate(0, 0, 1)
}

func (service *Scheduler) process(tm time.Time) {
	if !service.lastExecution.IsZero() && tm.After(service.nextExecution) {
		if err := service.insertData(); err != nil {
			service.log.Error().Err(err).Msgf("fail to scheduler insert data %+v", tm)
			return
		}
		service.log.Info().Msgf("행정동 insert date %+v", tm)
		service.lastExecution = tm
		service.executionTime()
	}
}

func (service *Scheduler) Run(ctx context.Context) <-chan error {
	c := make(chan error)
	m := time.NewTicker(60 * time.Second)

	go func() {
		defer close(c)
		defer m.Stop()

		for {
			select {
			case <-ctx.Done():
				service.log.Info().Msg("context canceled scheduler is stopping")
				break
			case tm := <-m.C:
				service.process(tm)
			}
		}
	}()

	return c
}

func NewScheduler(db *gorm.DB, log zerolog.Logger) *Scheduler {
	now := time.Now()
	return &Scheduler{db: db, log: log, processTime: now, lastExecution: now}
}
