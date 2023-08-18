package restapis

import (
	"area/application/service"
	"area/interfaces/restapis/context"
	"area/interfaces/restapis/sido"
	"area/interfaces/restapis/sigg"
	"common/response"
	"common/wrapper"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
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

func (a *TemplateService) Process(template context.AreaAPIService) error {
	ctx := template.GetContext()
	if ctx.Request().Method == http.MethodOptions {
		return ctx.NoContent(http.StatusOK)
	}

	a.logger.Debugf("header : %+v", ctx.Request().Header)

	request := template.GetRequestData()
	if request != nil {
		if err := ctx.Bind(request); err != nil {
			a.logger.WithError(err).Error("req parsing failed")
			template.Logger().Errorf("decode failed")
			return response.FailJSON(http.StatusBadRequest, "데이터 파싱 실패").Send(ctx)
		}
	}

	resp := template.Handler()
	if resp != nil {
		return resp.Send(ctx)
	}
	return nil
}

func (a *TemplateService) GetCities(c echo.Context) error {
	return a.Process(sido.NewCitesEndPoint(a.ParseContext(c)))
}

func (a *TemplateService) GetSigg(c echo.Context) error {
	return a.Process(sigg.NewSiggEndPoint(a.ParseContext(c)))
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

	route.GET("/cities", api.GetCities)
	route.GET("/sigg/:city_id", api.GetSigg)

	return api
}
