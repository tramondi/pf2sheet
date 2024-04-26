package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	common_dto "github.com/alionapermes/pf2sheet/internal/api/http/common/dto"
	"github.com/alionapermes/pf2sheet/internal/app/resource"
	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

func GetAllSheets(container resource.Container) echo.HandlerFunc {
	type dtoSheet struct {
		ID         int     `json:"id"`
		Ancestry   *string `json:"ancestry,omitempty"`
		Class      *string `json:"class,omitempty"`
		Background *string `json:"background,omitempty"`
		FullName   *string `json:"full_name,omitempty"`
		Level      *int16  `json:"level,omitempty"`
		HpCurrent  *int16  `json:"hp_current,omitempty"`
		HpMax      *int16  `json:"hp_max,omitempty"`
	}

	type data struct {
		Sheets []dtoSheet `json:"sheets"`
	}

	getAllSheets := container.Usecases.GetAllSheets

	return func(ctx echo.Context) error {
		id, ok := ctx.Get("player_id").(entity.PlayerID)
		if !ok {
			return ctx.NoContent(http.StatusInternalServerError)
		}

		sheets, err := getAllSheets.Execute(context.Background(), id)
		if err != nil {
			return ctx.NoContent(http.StatusInternalServerError)
		}

		sheetDTOs := make([]dtoSheet, 0, len(sheets))
		for _, sheet := range sheets {
			var ancestry *string
			var class *string

			if value := sheet.Ancestry; value != nil {
				ancestry = &value.Code
			}

			if value := sheet.Class; value != nil {
				class = &value.Code
			}

			sheetDTOs = append(sheetDTOs, dtoSheet{
				ID:         sheet.ID.Value(),
				Ancestry:   ancestry,
				Class:      class,
				Background: sheet.Background,
				FullName:   sheet.FullName,
				Level:      sheet.Level,
				HpCurrent:  sheet.HpCurrent,
				HpMax:      sheet.HpMax,
			})
		}

		response := &common_dto.Response{
			Data: &data{Sheets: sheetDTOs},
		}

		return ctx.JSON(http.StatusOK, response)
	}
}
