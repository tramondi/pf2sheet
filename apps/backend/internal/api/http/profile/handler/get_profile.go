package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	common_dto "github.com/alionapermes/pf2sheet/internal/api/http/common/dto"
	"github.com/alionapermes/pf2sheet/internal/app/resource"
	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

func GetProfile(container resource.Container) echo.HandlerFunc {
	type data struct {
		Login       string  `json:"login"`
		DisplayName *string `json:"display_name,omitempty"`
	}

	getProfile := container.Usecases.GetPlayer

	return func(ctx echo.Context) error {
		playerID, ok := ctx.Get("player_id").(entity.PlayerID)
		if !ok {
			ctx.Logger().Error("failed to get player_id from ctx")
			return ctx.NoContent(http.StatusInternalServerError)
		}

		player, err := getProfile.Execute(
			context.Background(),
			playerID,
		)
		if err != nil {
			ctx.Logger().Error("failed to execute getProfile usecase: %s", err)
			return ctx.NoContent(http.StatusInternalServerError)
		}

		response := &common_dto.Response{Data: data{
			Login:       player.Login,
			DisplayName: player.DisplayName,
		}}

		return ctx.JSON(http.StatusOK, response)
	}
}
