package restapis

import (
	"area/application/service"
	"area/interfaces/restapis/context"
	"common/wrapper"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"time"
)

type TemplateService struct {
	baseURI string
	*wrapper.EchoWrapper
	logger  *logrus.Entry
	service *service.Service
}

func (a *TemplateService) ParseContext(c echo.Context) *context.APIContext {
	return c.(*context.APIContext)
}

func (a *TemplateService) MiddlewareContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := &context.APIContext{
			Context: c,
			App:     a.service,
			Log:     a.logger,
		}
		ctx.Set("request_time", time.Now())
		err := next(ctx)
		return err
	}
}

func NewAPIService(port string, app *service.Service) *TemplateService {
	api := &TemplateService{
		baseURI:     "/api/v1/area",
		EchoWrapper: wrapper.NewEchoWrapper(port),
		logger:      logrus.WithField("component", "api-area"),
		service:     app,
	}

	route := api.SetGroup(api.baseURI, api.MiddlewareContext)

	route.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, "pong")
	})

	return api
}
