package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/alionapermes/pf2sheet/internal/api/http/common"
	common_dto "github.com/alionapermes/pf2sheet/internal/api/http/common/dto"
	"github.com/alionapermes/pf2sheet/internal/app/resource"
	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

func CreateSheet(container resource.Container) echo.HandlerFunc {
	type sheetDTO struct {
		AncestryID *int64  `json:"ancestry_id,omitempty"`
		ClassID    *int64  `json:"class_id,omitempty"`
		Background *string `json:"background,omitempty"`
		FullName   *string `json:"full_name,omitempty"`
		Level      *int16  `json:"level,omitempty"`
		HpCurrent  *int16  `json:"hp_current,omitempty"`
		HpMax      *int16  `json:"hp_max,omitempty"`
		Note       *string `json:"note,omitempty"`
	}

	type data struct {
		SheetID int `json:"sheet_id"`
	}

	createSheet := container.Usecases.CreateSheet

	return func(ctx echo.Context) error {
		playerID, ok := ctx.Get("player_id").(entity.PlayerID)
		if !ok {
			return ctx.NoContent(http.StatusInternalServerError)
		}

		var reqDTO sheetDTO

		if err := ctx.Bind(&reqDTO); err != nil {
			ctx.Logger().Errorf("failed to parse body: %s", err)
			return ctx.NoContent(http.StatusBadRequest)
		}

		var ancestry *entity.Ancestry
		var class *entity.Class

		if reqDTO.AncestryID != nil {
			ancestry = &entity.Ancestry{ID: entity.AncestryID(*reqDTO.AncestryID)}
		}

		if reqDTO.ClassID != nil {
			class = &entity.Class{ID: entity.ClassID(*reqDTO.ClassID)}
		}

		sheet := entity.Sheet{
			PlayerID:   playerID,
			Ancestry:   ancestry,
			Class:      class,
			Background: common.ClearTextPtr(reqDTO.Background),
			FullName:   common.ClearTextPtr(reqDTO.FullName),
			Note:       common.EscapeScriptsPtr(reqDTO.Note),
			Level:      reqDTO.Level,
			HpCurrent:  reqDTO.HpCurrent,
			HpMax:      reqDTO.HpMax,
		}

		id, err := createSheet.Execute(context.Background(), playerID, sheet)
		if err != nil {
			ctx.Logger().Errorf("failed to execute createSheet usecase: %s", err)
			return ctx.NoContent(http.StatusInternalServerError)
		}

		response := &common_dto.Response{
			Data: &data{SheetID: id.Value()},
		}

		return ctx.JSON(http.StatusOK, response)
	}
}
