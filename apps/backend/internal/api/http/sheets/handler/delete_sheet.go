package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/alionapermes/pf2sheet/internal/app/resource"
	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

func DeleteSheet(container resource.Container) echo.HandlerFunc {
	deleteSheet := container.Usecases.DeleteSheet

	return func(ctx echo.Context) error {
		playerID, ok := ctx.Get("player_id").(entity.PlayerID)
		if !ok {
			return ctx.NoContent(http.StatusInternalServerError)
		}

		rawSheetID := ctx.Param("id")

		sheetID, err := strconv.Atoi(rawSheetID)
		if err != nil {
			ctx.Logger().Errorf("rawSheetID: %v", rawSheetID)
			return ctx.NoContent(http.StatusBadRequest)
		}

		if err := deleteSheet.Execute(
			context.Background(),
			playerID,
			entity.SheetID(sheetID),
		); err != nil {
			return ctx.NoContent(http.StatusInternalServerError)
		}

		return ctx.NoContent(http.StatusOK)
	}
}
