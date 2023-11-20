package service

import (
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"wd_user/pkg/model"
)

type User struct {
	SignUp *SignUp
}

func NewUser(db *gorm.DB, log zerolog.Logger) *User {
	repo := model.NewMemberRepository(db)
	return &User{
		SignUp: NewSignUp(repo, log),
	}
}
