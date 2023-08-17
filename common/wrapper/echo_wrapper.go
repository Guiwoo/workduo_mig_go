package wrapper

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"time"
)

type EchoWrapper struct {
	core   *echo.Echo
	Logger *logrus.Entry
	listen string
}

func (e *EchoWrapper) Run(ctx context.Context) <-chan error {
	defer e.Logger.Warnf("echo server shutdown")
	errStream := make(chan error)
	go func() {
		defer close(errStream)
		if err := e.core.Start(e.listen); err != nil {
			errStream <- err
		}

		cx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		if err := e.core.Shutdown(cx); err != nil {
			e.Logger.WithError(err).Errorln("echo server shutdown failure")
			errStream <- err
			return
		}
		return
	}()

	go func() {
		<-ctx.Done()
		fmt.Println("echo server shutdown by parents context")
		e.Logger.Warnf("echo server shutdown by parents context")
		cx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		if err := e.core.Shutdown(cx); err != nil {
			e.Logger.WithError(err).Errorln("echo server shutdown failure")
			errStream <- err
			return
		}
		return
	}()
	return errStream
}

func enableLog(e *echo.Echo, logger *logrus.Logger) {
	e.Use(middleware.BodyDump(func(e echo.Context, reqBody []byte, resBody []byte) {
		reqTime := GetRequestTime(e)
		req := e.Request()

		headers := logrus.Fields{
			"CLIENT_IP": e.RealIP(),
		}
		logger.WithFields(headers).Infof("Request URI : %s"+
			"\n\t[Request  :client->server] ip:%s method:%s path:%s body:%s "+
			"\n\t[Response :server->client] status:%d latency:%s body:%s",
			req.RequestURI, e.RealIP(), req.Method, req.RequestURI, string(reqBody),
			e.Response().Status, time.Now().Sub(reqTime).String(), string(resBody),
		)
	}))
}

func GetRequestTime(c echo.Context) time.Time {
	if t := c.Get("request_time"); t != nil {
		return time.Now()
	} else {
		return t.(time.Time)
	}
}

func NewEchoWrapper(listen string) *EchoWrapper {
	e := &EchoWrapper{
		core:   echo.New(),
		Logger: logrus.WithField("component", "echo"),
		listen: listen,
	}
	enableLog(e.core, logrus.StandardLogger())
	e.Logger = logrus.WithField("component", "area-server")
	e.core.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	return e
}
