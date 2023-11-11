package model

import (
	"context"
	"gorm.io/gorm"
)

const SiggAreaTableName = "sigg_area"

type SiggArea struct {
	AreaId   string   `gorm:"column:area_id;primaryKey"`
	CityId   string   `gorm:"column:city_id"`
	AmdCd    string   `gorm:"column:amd_cd"`
	SiggName string   `gorm:"column:siggnm"`
	Sido     SidoArea `gorm:"foreignKey:city_id;references:city_id"`
}

func (s *SiggArea) TableName() string {
	return SiggAreaTableName
}

func (s *SiggArea) GetAreaByCity(ctx context.Context, db *gorm.DB, id string) ([]SiggArea, error) {
	var result []SiggArea
	err := db.WithContext(ctx).Model(&SiggArea{}).Where("city_id = ?", id).Find(&result).Error
	return result, err
}
