package service

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"wd_common/wd_response"
	"wd_user/pkg/model"
)

type Logout struct {
	repo model.MemberRepository
	log  zerolog.Logger
}

func (service *Logout) Handle(ctx echo.Context) error {
	// logout 처리
	return wd_response.SuccessJSON(ctx, nil)
}

func NewLogout(repo model.MemberRepository, log zerolog.Logger) *Logout {
	return &Logout{
		repo: repo,
		log:  log,
	}
}
