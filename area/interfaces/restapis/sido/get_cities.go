package sido

import (
	"area/interfaces/restapis/context"
	"common/response"
	"github.com/sirupsen/logrus"
	"net/http"
)

type CitesEndPoint struct {
	*context.APIContext
	logger *logrus.Entry
}

func (c *CitesEndPoint) GetRequestData() interface{} {
	return nil
}

func (c *CitesEndPoint) ApiName() string {
	return "api-sido-get-cities"
}

func (c *CitesEndPoint) Handler() *response.WorkDuoResponse {
	sido, err := c.App.Area.GetAllCity(c.Request().Context())
	if err != nil {
		c.logger.WithError(err).Error("get cities failed")
		return response.FailJSON(http.StatusInternalServerError, "시도 정보를 가져오는데 실패했습니다.")
	}
	return response.SuccessJSON(sido)
}

func (c *CitesEndPoint) IsRequiredAuth() bool {
	return false
}

func (c *CitesEndPoint) Logger() *logrus.Entry {
	return c.logger
}

var _ context.AreaAPIService = (*CitesEndPoint)(nil)

func NewCitesEndPoint(ctx *context.APIContext) *CitesEndPoint {
	return &CitesEndPoint{
		APIContext: ctx,
		logger:     ctx.Log,
	}
}
