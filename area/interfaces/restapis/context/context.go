package context

import (
	"area/application/service"
	"context"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

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
