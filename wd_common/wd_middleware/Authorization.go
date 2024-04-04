package wd_middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"wd_common/wd_response"
	"wd_common/wd_token"
)

func Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		token := c.Request().Header.Get("Authorization")
		auth, err := wd_token.Parse(token)
		if err != nil {
			return wd_response.FailJSON(c, http.StatusUnauthorized, "Token is not valid")
		}

		c.Set("user", auth)

		return next(c)
	}
}
