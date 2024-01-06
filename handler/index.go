package handler

import (
	"be/web"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	return render(c, web.Index())
}
