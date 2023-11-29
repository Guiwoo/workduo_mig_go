package service

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"net/http"
	"regexp"
	"time"
	"wd_common/wd_response"
	"wd_user/pkg/model"
)

const FailSignUp = "회원가입에 실패했습니다."

type signUpRequest struct {
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Phone     string   `json:"phone"`
	Password  string   `json:"password"`
	Areas     []string `json:"areas"`
	Exercises []string `json:"exercises"`
}

type SignUp struct {
	repo model.MemberRepository
	log  zerolog.Logger
	req  signUpRequest
}

func (service *SignUp) validatePassword() bool {
	pattern := `.*[!@#$%^&*()_+{}\[\]:;<>,.?~\\/\\-].*`
	match, err := regexp.Match(pattern, []byte(service.req.Password))
	if err != nil {
		fmt.Println("Error while matching password pattern:", err)
		return false
	}

	return match
}

func (service *SignUp) validateEmail() bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if ok, err := regexp.MatchString(pattern, service.req.Email); !ok || err != nil {
		return false
	}
	return true
}

func (service *SignUp) validateRequest() error {
	if service.req.Name == "" {
		return ErrRequiredName
	}
	if service.validateEmail() == false {
		return ErrWrongFormatEmail
	}
	if service.validatePassword() == false {
		return ErrWrongFormatPassword
	}
	if len(service.req.Areas) < 1 {
		return ErrEmptyArea
	}
	if len(service.req.Exercises) < 1 {
		return ErrEmptyExercise
	}
	return nil
}

func (service *SignUp) Handle(ctx echo.Context) error {
	if err := ctx.Bind(&service.req); err != nil {
		service.log.Error().Err(err).Msgf("fail to bind req %+v", service.req)
		return wd_response.FailJSON(ctx, http.StatusBadRequest, WrongBodyForm)
	}
	if err := service.validateRequest(); err != nil {
		service.log.Error().Err(err).Msgf("fail to validate request %+v", service.req)
		return wd_response.FailJSON(ctx, http.StatusBadRequest, FailSignUp)
	}

	tb := &model.Member{
		MemberID:    IDGenerator("MB"),
		Name:        service.req.Name,
		Email:       service.req.Email,
		Nickname:    time.Now().String(),
		PhoneNumber: service.req.Phone,
		Password:    service.req.Password,
	}

	if err := tb.HashPassword(); err != nil {
		service.log.Error().Err(err).Msgf("fail to hash password %+v", tb)
		return wd_response.FailJSON(ctx, http.StatusInternalServerError, FailSignUp)
	}

	if err := service.repo.Create(ctx, tb); err != nil {
		service.log.Error().Err(err).Msgf("fail to create member %+v", tb)
		return wd_response.FailJSON(ctx, http.StatusInternalServerError, FailSignUp)
	}

	return wd_response.SuccessJSON(ctx, nil)
}

func NewSignUp(repo model.MemberRepository, log zerolog.Logger) *SignUp {
	return &SignUp{repo: repo, log: log}
}
