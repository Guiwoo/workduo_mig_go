package service

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"mime/multipart"
	"net/http"
	"wd_common/wd_response"
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

func (service *Update) Handle(ctx echo.Context) error {
	if err := service.Bind(ctx); err != nil {
		service.log.Err(err).Msgf("fail to bind request %+v", service.req)
		return wd_response.FailJSON(ctx, http.StatusBadRequest, UpdateError)
	}

	return wd_response.SuccessJSON(ctx, nil)
}

func NewUpdate(repo model.MemberRepository, log zerolog.Logger) *Update {
	return &Update{
		repo: repo,
		log:  log,
	}
}
