package model

import (
	"context"
	"gorm.io/gorm"
)

const SidoTableName = "sido_area"

type SidoArea struct {
	CityId   string `gorm:"column:city_id;primaryKey"`
	AdmCd    string `gorm:"column:adm_cd"`
	SidoName string `gorm:"column:sidonm"`
}

func (s *SidoArea) TableName() string {
	return SidoTableName
}

func (s *SidoArea) GetAllCity(ctx context.Context, db *gorm.DB) ([]SidoArea, error) {
	var result []SidoArea
	err := db.WithContext(ctx).Model(&SidoArea{}).Find(&result).Error
	return result, err
}
