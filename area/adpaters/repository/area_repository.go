package repository

import (
	"area/adpaters"
	"area/domain"
	"area/infrastructure/persistence"
	"context"
	"gorm.io/gorm"
)

type AreaDBRepository struct {
	db *gorm.DB
}

func (a *AreaDBRepository) GetAllCity(ctx context.Context) ([]domain.Sido, error) {
	var data persistence.SidoArea
	result, err := data.GetAllCity(ctx, a.db)
	if err != nil {
		return nil, err
	}
	sido := make([]domain.Sido, len(result))
	for i := range sido {
		sido[i] = domain.Sido{
			SidoID:   result[i].CityId,
			SidoName: result[i].SidoName,
		}
	}
	return sido, nil
}

func (a *AreaDBRepository) GetAreaByCity(ctx context.Context, cityId string) ([]domain.Sigg, error) {
	var data persistence.SiggArea
	result, err := data.GetAreaByCity(ctx, a.db, cityId)
	if err != nil {
		return nil, err
	}
	sigg := make([]domain.Sigg, len(result))
	for i := range sigg {
		sigg[i] = domain.Sigg{
			SiggID:   result[i].AreaId,
			SiggName: result[i].SiggName,
			SidoID:   result[i].CityId,
			SidoName: result[i].Sido.SidoName,
		}
	}
	return sigg, nil
}

var _ adpaters.AreaRepository = (*AreaDBRepository)(nil)

func NewAreaDBRepository(db *gorm.DB) adpaters.AreaRepository {
	return &AreaDBRepository{
		db: db,
	}
}
