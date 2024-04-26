package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	common_dto "github.com/alionapermes/pf2sheet/internal/api/http/common/dto"
	"github.com/alionapermes/pf2sheet/internal/api/http/knowledge/dto"
	"github.com/alionapermes/pf2sheet/internal/app/resource"
)

func GetAncestries(container resource.Container) echo.HandlerFunc {
	type data struct {
		Ancestries []dto.Ancestry `json:"ancestries"`
	}

	getAllAncestries := container.Usecases.GetAllAncestries

	return func(ctx echo.Context) error {
		ancestries, err := getAllAncestries.Execute(context.TODO())
		if err != nil {
			return ctx.String(http.StatusInternalServerError, "internal error")
		}

		DTOs := make([]dto.Ancestry, 0, len(ancestries))
		for _, ancestry := range ancestries {
			DTOs = append(DTOs, dto.Ancestry{
				ID:    int64(ancestry.ID),
				Title: ancestry.Title,
			})
		}

		response := &common_dto.Response{
			Data: &data{Ancestries: DTOs},
		}

		return ctx.JSON(http.StatusOK, response)
	}
}
