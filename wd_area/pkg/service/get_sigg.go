package service

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"net/http"
	"wd_area/pkg/model"
	"wd_common/wd_response"
)

const FailGetSiggService = "시군구 조회를 실패 했습니다."

type responseGetSiggService struct {
	SiggID   string
	SiggName string
	SidoID   string
	SidoName string
}
type requestGetSigg struct {
	ID string `param:"city_id"`
}

type GetSigg struct {
	db  *gorm.DB
	log zerolog.Logger
	req requestGetSigg
}

func (service *GetSigg) Handle(ctx echo.Context) error {
	if err := ctx.Bind(&service.req); err != nil {
		return wd_response.FailJSON(ctx, http.StatusBadRequest, FailGetSiggService)
	}

	tb := model.SiggArea{}
	data, err := tb.GetAreaByCity(ctx.Request().Context(), service.db, service.req.ID)
	if err != nil {
		service.log.Err(err).Msg("fail to get sigg data")
		return wd_response.FailJSON(ctx, http.StatusInternalServerError, FailGetSiggService)
	}

	sigg := make([]responseGetSiggService, len(data))
	for i := range sigg {
		sigg[i] = responseGetSiggService{
			SiggID:   data[i].AreaId,
			SiggName: data[i].SiggName,
			SidoID:   data[i].CityId,
			SidoName: data[i].Sido.SidoName,
		}
	}

	return wd_response.SuccessJSON(ctx, sigg)
}

func NewGetSigg(db *gorm.DB, log zerolog.Logger) *GetSigg {
	return &GetSigg{
		db:  db,
		log: log,
	}
}
