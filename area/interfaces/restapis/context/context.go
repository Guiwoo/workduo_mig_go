package context

import (
	"area/application/service"
	"common/response"
	"context"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type AreaAPIService interface {
	GetRequestData() interface{}
	ApiName() string
	Handler() *response.WorkDuoResponse
	IsRequiredAuth() bool
	GetContext() echo.Context
	Logger() *logrus.Entry
}

type APIContext struct {
	echo.Context
	App *service.Service
	Log *logrus.Entry
}

func (c *APIContext) GetContext() echo.Context {
	return c.Context
}
func (c *APIContext) GetRequestContext() context.Context {
	return c.Request().Context()
}
