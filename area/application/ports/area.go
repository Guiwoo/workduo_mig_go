package ports

import (
	"area/adpaters"
	"area/adpaters/repository"
	"gorm.io/gorm"
)

type AreaHandler struct {
	repo adpaters.AreaRepository
}

func NewAreaHandler(db *gorm.DB) *AreaHandler {
	return &AreaHandler{repository.NewAreaDBRepository(db)}
}
