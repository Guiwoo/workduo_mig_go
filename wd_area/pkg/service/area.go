package service

import (
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Area struct {
	GetCities *GetCities
	GetSigg   *GetSigg
}

func New(db *gorm.DB, log zerolog.Logger) *Area {
	return &Area{
		GetCities: NewGetCities(db, log),
		GetSigg:   NewGetSigg(db, log),
	}
}
