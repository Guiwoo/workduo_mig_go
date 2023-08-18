package service

import (
	"area/application/ports"
	"gorm.io/gorm"
)

type Service struct {
	Area *ports.AreaHandler
}

func NewService(areaDB *gorm.DB) *Service {
	return &Service{
		Area: ports.NewAreaHandler(areaDB),
	}
}
