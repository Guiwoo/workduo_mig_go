package service

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"net/http"
	"wd_area/pkg/model"
	"wd_common/response"
)

const FailGetSiggService = "시군구 조회를 실패 했습니다."

type responseGetSiggService struct {
	SiggID   string
	SiggName string
	SidoID   string
	SidoName string
}
type requestGetSiggService struct {
	ID string `param:"city_id"`
}

type GetSiggService struct {
	db  *gorm.DB
	log zerolog.Logger
	req requestGetSiggService
}

func (service *GetSiggService) Handle(ctx echo.Context) error {
	if err := ctx.Bind(&service.req); err != nil {
		return response.FailJSON(ctx, http.StatusBadRequest, FailGetSiggService)
	}

	tb := model.SiggArea{}
	data, err := tb.GetAreaByCity(ctx.Request().Context(), service.db, service.req.ID)
	if err != nil {
		service.log.Err(err).Msg("fail to get sigg data")
		return response.FailJSON(ctx, http.StatusInternalServerError, FailGetSiggService)
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

	return response.SuccessJSON(ctx, sigg)
}

func NewGetSiggService(db *gorm.DB, log zerolog.Logger) *GetSiggService {
	return &GetSiggService{
		db:  db,
		log: log,
	}
}
