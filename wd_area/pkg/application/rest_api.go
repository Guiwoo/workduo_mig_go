package application

import (
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"wd_area/pkg/service"
	"wd_common/wd_wrapper"
)

const URL = "/area/api/v1.0"

type AreaRestAPI struct {
	*wd_wrapper.EchoWrapper
	service *service.Area
}

func (api *AreaRestAPI) route() {
	core := api.SetGroup(URL)
	core.GET("/cities", api.service.GetCities.Handle)
	core.GET("/sigg", api.service.GetSigg.Handle)
}

func NewAreaRestAPI(db *gorm.DB, port string, log zerolog.Logger) *AreaRestAPI {
	api := &AreaRestAPI{
		EchoWrapper: wd_wrapper.NewEcho(port),
		service:     service.New(db, log),
	}

	api.route()

	return api
}
