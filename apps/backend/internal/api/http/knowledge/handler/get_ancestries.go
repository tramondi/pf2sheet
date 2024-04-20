package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/alionapermes/pf2sheet/internal/api/http/knowledge/dto"
	"github.com/alionapermes/pf2sheet/internal/app/resource"
)

func GetAncestries(usecases *resource.Usecases) echo.HandlerFunc {
	type result struct {
		Ancestries []dto.Ancestry `json:"ancestries"`
	}

	return func(ctx echo.Context) error {
		ancestries, err := usecases.GetAllAncestries.Execute(context.TODO())
		if err != nil {
			return ctx.String(http.StatusInternalServerError, "kek")
		}

		DTOs := make([]dto.Ancestry, 0, len(ancestries))
		for _, ancestry := range ancestries {
			DTOs = append(DTOs, dto.Ancestry{
				Code:  ancestry.Code,
				Title: ancestry.Title,
			})
		}

		response := &dto.Response[result]{
			Data: &result{Ancestries: DTOs},
		}

		return ctx.JSON(http.StatusOK, response)
	}
}
