package service

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"mime/multipart"
	"net/http"
	"regexp"
	"wd_common/wd_response"
	"wd_common/wd_token"
	"wd_user/pkg/model"
)

const UpdateError = "멤버 업데이트에 실패했습니다."

type Update struct {
	repo model.MemberRepository
	log  zerolog.Logger
	req  UpdateRequest
}

type UpdateRequest struct {
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Nickname  string `json:"nickname"`
	StatusMsg string `json:"status_msg"`
	File      *multipart.FileHeader
}

func (service *Update) Bind(ctx echo.Context) error {
	service.req.Name = ctx.FormValue("name")
	service.req.Phone = ctx.FormValue("phone")
	service.req.Nickname = ctx.FormValue("nickname")
	service.req.StatusMsg = ctx.FormValue("status_msg")
	multiFile, err := ctx.FormFile("file")
	if err != nil {
		return err
	}

	service.req.File = multiFile
	return nil
}

func (service *Update) validate() error {
	if service.req.Name == "" {
		return fmt.Errorf("name is required")
	}

	pattern := `^\d{3}-\d{4}-\d{4}$`
	_pattern := regexp.MustCompile(pattern)
	if ok := _pattern.MatchString(service.req.Phone); ok == false {
		return fmt.Errorf("phone format is wrong")
	}

	if service.req.Nickname == "" {
		return fmt.Errorf("nick name is required")
	}

	if service.req.StatusMsg == "" {
		return fmt.Errorf("nickname status is required")
	}

	return nil
}

func (service *Update) Handle(ctx echo.Context) error {
	user, ok := ctx.Get("user").(*wd_token.Token)
	if ok == false {
		service.log.Err(fmt.Errorf("fail to get user"))
		return wd_response.FailJSON(ctx, http.StatusUnauthorized, "유저를 찾을수 없습니다.")
	}

	if err := service.Bind(ctx); err != nil {
		service.log.Err(err).Msgf("fail to bind request %+v", service.req)
		return wd_response.FailJSON(ctx, http.StatusBadRequest, UpdateError)
	}

	if err := service.validate(); err != nil {
		service.log.Err(err).Msgf("fail to validate service %+v", service.req)
		return wd_response.FailJSON(ctx, http.StatusBadRequest, err.Error())
	}

	m := &model.Member{
		MemberID: user.MemberID,
		Name:     service.req.Name,
		Nickname: service.req.Nickname,
		//todo aws address
		ProfileIMG:   "",
		MemberStatus: service.req.StatusMsg,
	}

	if err := service.repo.Update(ctx, m); err != nil {
		service.log.Err(err).Msgf("fail to update member %+v", err)
		return wd_response.FailJSON(ctx, http.StatusInternalServerError, UpdateError)
	}

	return wd_response.SuccessJSON(ctx, nil)
}

func NewUpdate(repo model.MemberRepository, log zerolog.Logger) *Update {
	return &Update{
		repo: repo,
		log:  log,
	}
}
