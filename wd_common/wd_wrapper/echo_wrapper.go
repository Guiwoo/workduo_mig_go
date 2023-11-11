package wd_wrapper

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"time"
)

type EchoWrapper struct {
	core   *echo.Echo
	listen string
}

func (e *EchoWrapper) Run(ctx context.Context) <-chan error {
	errStream := make(chan error)
	go func() {
		defer close(errStream)
		if err := e.core.Start(e.listen); err != nil {
			errStream <- err
		}

		cx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		if err := e.core.Shutdown(cx); err != nil {
			errStream <- err
			return
		}
		return
	}()

	go func() {
		<-ctx.Done()
		fmt.Println("echo server shutdown by parents context")
		cx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		if err := e.core.Shutdown(cx); err != nil {
			errStream <- err
			return
		}
		return
	}()
	return errStream
}

func (e *EchoWrapper) SetGroup(prefix string, m ...echo.MiddlewareFunc) *echo.Group {
	return e.core.Group(prefix, m...)
}

func NewEcho(listen string) *EchoWrapper {
	e := &EchoWrapper{
		core:   echo.New(),
		listen: listen,
	}

	return e
}
