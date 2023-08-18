package adpaters

import (
	"area/domain"
	"context"
)

type AreaRepository interface {
	GetAllCity(ctx context.Context) ([]domain.Sido, error)
	GetAreaByCity(ctx context.Context, cityId string) ([]domain.Sigg, error)
}
