package service

import (
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type AreaService struct {
	GetCities *GetCitiesService
	GetSigg   *GetSiggService
}

func New(db *gorm.DB, log zerolog.Logger) *AreaService {
	return &AreaService{
		GetCities: NewGetCitiesService(db, log),
		GetSigg:   NewGetSiggService(db, log),
	}
}
