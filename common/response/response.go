package response

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type WorkDuoResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func (w *WorkDuoResponse) Send(c echo.Context) error {
	return c.JSON(http.StatusOK, w)
}

func FailJSON(code int, msg string) *WorkDuoResponse {
	return &WorkDuoResponse{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}
func SuccessJSON(data interface{}) *WorkDuoResponse {
	return &WorkDuoResponse{
		Code: http.StatusOK,
		Msg:  "success",
		Data: data,
	}
}
