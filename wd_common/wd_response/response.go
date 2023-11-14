package wd_response

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type WorkDuoResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func FailJSON(ctx echo.Context, code int, msg string) error {
	return ctx.JSON(http.StatusOK, &WorkDuoResponse{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
func SuccessJSON(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusOK, &WorkDuoResponse{
		Code: http.StatusOK,
		Msg:  "success",
		Data: data,
	})
}
