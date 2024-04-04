package application

import (
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"wd_common/wd_middleware"
	"wd_common/wd_wrapper"
	"wd_user/pkg/service"
)

const URL = "/user/api/v1.0"

type UserRestAPI struct {
	*wd_wrapper.EchoWrapper
	service *service.Member
}

func (api *UserRestAPI) route() {
	core := api.SetGroup(URL, wd_middleware.Authorization)

	core.POST("/member", api.service.SignUp.Handle)
	core.POST("/login", api.service.Login.Handle)
	core.POST("/logout", api.service.Logout.Handle)
	core.PATCH("/member", api.service.Update.Handle)
}

func NewUserRestAPI(db *gorm.DB, port string, log zerolog.Logger) *UserRestAPI {
	api := &UserRestAPI{
		EchoWrapper: wd_wrapper.NewEcho(port),
		service:     service.NewMember(db, log),
	}

	api.route()
	return api
}
