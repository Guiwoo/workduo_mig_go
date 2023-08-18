package area

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
