package area

const SidoAreaTableName = "sido_area"

type SidoArea struct {
	CityId   string `gorm:"column:city_id;primaryKey"`
	AdmCd    string `gorm:"column:adm_cd"`
	SidoName string `gorm:"column:sidonm"`
}

func (s *SidoArea) TableName() string {
	return SidoAreaTableName
}
