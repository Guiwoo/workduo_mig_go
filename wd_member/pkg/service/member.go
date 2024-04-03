package service

import (
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"wd_user/pkg/model"
)

type Member struct {
	SignUp *SignUp
	Login  *Login
	Logout *Logout
	Update *Update
}

func NewMember(db *gorm.DB, log zerolog.Logger) *Member {
	repo := model.NewMemberRepository(db)
	return &Member{
		SignUp: NewSignUp(repo, log),
		Login:  NewLogin(repo, log),
		Logout: NewLogout(repo, log),
		Update: NewUpdate(repo, log),
	}
}
