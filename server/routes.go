package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) addRoutes(e *echo.Echo) {

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
}
