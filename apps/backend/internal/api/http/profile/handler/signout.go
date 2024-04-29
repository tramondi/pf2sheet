package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/alionapermes/pf2sheet/internal/app/resource"
	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

func Signout(container resource.Container) echo.HandlerFunc {
	signout := container.Usecases.Signout

	return func(ctx echo.Context) error {
		playerID, ok := ctx.Get("player_id").(entity.PlayerID)
		if !ok {
			ctx.Logger().Error("failed to get player_id from ctx")
			return ctx.NoContent(http.StatusInternalServerError)
		}

		token, ok := ctx.Get("session_token").(string)
		if !ok {
			ctx.Logger().Error("failed to get session_token from ctx")
			return ctx.NoContent(http.StatusInternalServerError)
		}

		if err := signout.Execute(
			context.Background(),
			playerID,
			token,
		); err != nil {
			ctx.Logger().Errorf("failed to execute signout usecase: %s", err)
			return ctx.NoContent(http.StatusInternalServerError)
		}

		return unsetSessionCookie(ctx).
			NoContent(http.StatusOK)
	}
}
