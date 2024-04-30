package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/alionapermes/pf2sheet/internal/app/resource"
	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

func DeleteProfile(container resource.Container) echo.HandlerFunc {
	deleteProfile := container.Usecases.DeletePlayer

	return func(ctx echo.Context) error {
		playerID, ok := ctx.Get("player_id").(entity.PlayerID)
		if !ok {
			ctx.Logger().Error("failed to get player_id from ctx")
			return ctx.NoContent(http.StatusInternalServerError)
		}

		if err := deleteProfile.Execute(
			context.Background(),
			playerID,
		); err != nil {
			ctx.Logger().Error("failed to execute deleteProfile usecase: %s", err)
			return ctx.NoContent(http.StatusInternalServerError)
		}

		return ctx.NoContent(http.StatusOK)
	}
}
