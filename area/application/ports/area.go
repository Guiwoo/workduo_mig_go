package ports

import (
	"area/adpaters"
	"area/adpaters/repository"
	"area/domain"
	"context"
	"gorm.io/gorm"
)

type AreaHandler struct {
	repo adpaters.AreaRepository
}

func (a *AreaHandler) GetAllCity(ctx context.Context) ([]domain.Sido, error) {
	return a.repo.GetAllCity(ctx)
}
func (a *AreaHandler) GetAreaByCity(ctx context.Context, city string) ([]domain.Sigg, error) {
	return a.repo.GetAreaByCity(ctx, city)
}

func NewAreaHandler(db *gorm.DB) *AreaHandler {
	return &AreaHandler{repository.NewAreaDBRepository(db)}
}
