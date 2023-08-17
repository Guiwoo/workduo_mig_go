package area

const SiggAreaTableName = "sigg_area"

type SiggArea struct {
	AreaId   string `gorm:"column:area_id;primaryKey"`
	CityId   string `gorm:"column:city_id"`
	AmdCd    string `gorm:"column:amd_cd"`
	SiggName string `gorm:"column:siggnm"`
}

func (s *SiggArea) TableName() string {
	return SiggAreaTableName
}
