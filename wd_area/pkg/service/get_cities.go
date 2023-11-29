package service

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"net/http"
	"wd_area/pkg/model"
	"wd_common/wd_response"
)

const FailGetCitiesService = "도시 조회를 실패했습니다."

type responseGetCities struct {
	SidoID   string
	SidoName string
}

type GetCities struct {
	db  *gorm.DB
	log zerolog.Logger
}

func (service GetCities) Handle(ctx echo.Context) error {
	tb := &model.SidoArea{}
	data, err := tb.GetAllCity(ctx.Request().Context(), service.db)
	if err != nil {
		service.log.Err(err).Msg("fail to get db data")
		return wd_response.FailJSON(ctx, http.StatusInternalServerError, FailGetCitiesService)
	}

	sido := make([]responseGetCities, len(data))
	for i := range sido {
		sido[i] = responseGetCities{
			SidoID:   data[i].CityId,
			SidoName: data[i].SidoName,
		}
	}

	return wd_response.SuccessJSON(ctx, sido)
}

func NewGetCities(db *gorm.DB, log zerolog.Logger) *GetCities {
	return &GetCities{db: db, log: log}
}
