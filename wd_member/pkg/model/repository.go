package model

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type MemberRepository interface {
	Create(ctx echo.Context, member *Member) error
}

type memberDB struct {
	db *gorm.DB
}

func (m *memberDB) Create(ctx echo.Context, member *Member) error {
	return member.create(ctx.Request().Context(), m.db)
}

var _ MemberRepository = (*memberDB)(nil)

func NewMemberRepository(db *gorm.DB) MemberRepository {
	return &memberDB{db: db}
}
