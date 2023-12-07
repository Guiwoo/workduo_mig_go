package service

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"net/http"
	"regexp"
	"wd_common/wd_response"
	"wd_common/wd_token"
	"wd_user/pkg/model"
)

const FailLogin = "로그인에 실패했습니다."

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Login struct {
	repo model.MemberRepository
	log  zerolog.Logger
	req  loginRequest
}

func (service *Login) validateEmail() bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if ok, err := regexp.MatchString(pattern, service.req.Email); !ok || err != nil {
		return false
	}
	return true
}

func (service *Login) validateRequest() error {
	if service.validateEmail() == false {
		return ErrWrongFormatEmail
	}
	return nil
}

func (service *Login) getHashedPassword() (string, error) {
	m := &model.Member{Password: service.req.Password}
	if err := m.HashPassword(); err != nil {
		service.log.Error().Err(err).Msgf("fail to hash password %+v", service.req)
		return "", err
	}
	return m.Password, nil
}

func (service *Login) Handle(ctx echo.Context) error {
	if err := ctx.Bind(&service.req); err != nil {
		service.log.Error().Err(err).Msgf("fail to bind req %+v", service.req)
		return wd_response.FailJSON(ctx, http.StatusBadRequest, WrongBodyForm)
	}

	if err := service.validateRequest(); err != nil {
		service.log.Error().Err(err).Msgf("fail to validate request %+v", service.req)
		return wd_response.FailJSON(ctx, http.StatusBadRequest, FailLogin)
	}

	member, err := service.repo.Find(ctx, service.req.Email)
	if err != nil {
		service.log.Error().Err(err).Msgf("fail to get member %+v", service.req)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return wd_response.FailJSON(ctx, ErrCodeEmail, FailLogin)
		}
		return wd_response.FailJSON(ctx, http.StatusInternalServerError, FailLogin)
	}

	hashedPassword, err := service.getHashedPassword()
	if err != nil {
		service.log.Error().Err(err).Msgf("fail to hashing password %+v", service.req)
		return wd_response.FailJSON(ctx, http.StatusBadRequest, FailLogin)
	}

	if hashedPassword != member.Password {
		return wd_response.FailJSON(ctx, ErrCodePassword, FailLogin)
	}

	tkn, err := wd_token.Generate(member.MemberID, member.Email)
	if err != nil {
		return wd_response.FailJSON(ctx, http.StatusInternalServerError, FailLogin)
	}

	return wd_response.SuccessJSON(ctx, struct {
		Token string `json:"token"`
	}{Token: tkn})
}

func NewLogin(repo model.MemberRepository, log zerolog.Logger) *Login {
	return &Login{repo: repo, log: log}
}
