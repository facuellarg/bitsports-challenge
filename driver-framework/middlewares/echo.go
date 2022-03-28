package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func EchoWrapper(handler http.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		handler.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
