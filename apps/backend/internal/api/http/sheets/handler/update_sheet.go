package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/alionapermes/pf2sheet/internal/app/resource"
	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

func UpdateSheet(container resource.Container) echo.HandlerFunc {
	type reqParams struct {
		FullName   *string `json:"full_name"`
		Level      *int16  `json:"level"`
		AncestryID *int64  `json:"ancestry_id"`
		ClassID    *int64  `json:"class_id"`
		Background *string `json:"background"`
		HpCurrent  *int16  `json:"hp_current"`
		HpMax      *int16  `json:"hp_max"`
	}

	updateSheet := container.Usecases.UpdateSheet

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

		var sheetDTO reqParams
		if err := ctx.Bind(&sheetDTO); err != nil {
			ctx.Logger().Errorf("failed to bind request dto: %s", err.Error())
			return ctx.NoContent(http.StatusInternalServerError)
		}

		var ancestry *entity.Ancestry
		var class *entity.Class

		if sheetDTO.AncestryID != nil {
			ancestry = &entity.Ancestry{ID: entity.AncestryID(*sheetDTO.AncestryID)}
		}

		if sheetDTO.ClassID != nil {
			class = &entity.Class{ID: entity.ClassID(*sheetDTO.ClassID)}
		}

		sheet := entity.Sheet{
			ID:         entity.SheetID(sheetID),
			PlayerID:   playerID,
			Ancestry:   ancestry,
			Class:      class,
			FullName:   sheetDTO.FullName,
			Level:      sheetDTO.Level,
			Background: sheetDTO.Background,
			HpCurrent:  sheetDTO.HpCurrent,
			HpMax:      sheetDTO.HpMax,
		}

		if err := updateSheet.Execute(context.Background(), sheet); err != nil {
			return ctx.NoContent(http.StatusInternalServerError)
		}

		return ctx.NoContent(http.StatusOK)
	}
}
