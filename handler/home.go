package handler

import (
	"be/web"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	return render(c, web.Home())
}
