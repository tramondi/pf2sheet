package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/alionapermes/pf2sheet/internal/app/resource"
	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

func UpdateProfile(container resource.Container) echo.HandlerFunc {
	type playerDTO struct {
		DisplayName *string `json:"display_name"`
	}

	updProfile := container.Usecases.UpdatePlayer

	return func(ctx echo.Context) error {
		playerID, ok := ctx.Get("player_id").(entity.PlayerID)
		if !ok {
			ctx.Logger().Error("failed to get player_id from ctx")
			return ctx.NoContent(http.StatusInternalServerError)
		}

		var dto playerDTO
		if err := ctx.Bind(&dto); err != nil {
			return ctx.NoContent(http.StatusBadRequest)
		}

		player := entity.Player{
			ID:          playerID,
			DisplayName: dto.DisplayName,
		}

		if err := updProfile.Execute(context.Background(), player); err != nil {
			ctx.Logger().Errorf("failed to execute updProfile usecase: %s", err)
			return ctx.NoContent(http.StatusInternalServerError)
		}

		return ctx.NoContent(http.StatusOK)
	}
}
