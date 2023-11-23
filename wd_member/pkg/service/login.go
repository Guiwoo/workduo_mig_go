package service

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"wd_user/pkg/model"
)

type loginRequest struct {
}

type Login struct {
	repo model.MemberRepository
	log  zerolog.Logger
	req  loginRequest
}

func (service *Login) Handle(ctx echo.Context) error {
	// [API 명세서] https://alive-tern-b83.notion.site/f4d237db90084099a0db9eebb2fea03e
	// [Spring-boot] https://github.com/Guiwoo/WorkDuo_dev
	return nil
}

func NewLogin(repo model.MemberRepository, log zerolog.Logger) *Login {
	return &Login{repo: repo, log: log}
}
