package sigg

import (
	"area/interfaces/restapis/context"
	"common/response"
	"github.com/sirupsen/logrus"
	"net/http"
)

type SiggEndPointRequest struct {
	ID string `param:"city_id"`
}

type SiggEndPoint struct {
	*context.APIContext
	req    SiggEndPointRequest
	logger *logrus.Entry
}

func (s *SiggEndPoint) GetRequestData() interface{} {
	return &s.req
}

func (s *SiggEndPoint) ApiName() string {
	return "api-sigg-get-sigg"
}

func (s *SiggEndPoint) Handler() *response.WorkDuoResponse {
	result, err := s.App.Area.GetAreaByCity(s.GetRequestContext(), s.req.ID)
	if err != nil {
		logrus.WithError(err).Error("failed to get sigg")
		return response.FailJSON(http.StatusInternalServerError, "시군구 조회 실패")
	}
	return response.SuccessJSON(result)
}

func (s *SiggEndPoint) IsRequiredAuth() bool {
	return false
}

func (s *SiggEndPoint) Logger() *logrus.Entry {
	return s.logger
}

var _ context.AreaAPIService = (*SiggEndPoint)(nil)

func NewSiggEndPoint(ctx *context.APIContext) *SiggEndPoint {
	return &SiggEndPoint{
		APIContext: ctx,
		logger:     ctx.Log,
	}
}
