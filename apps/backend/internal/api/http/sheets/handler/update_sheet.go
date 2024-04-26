package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/alionapermes/pf2sheet/internal/app/resource"
)

func UpdateSheet(container resource.Container) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusNotImplemented)
	}
}
