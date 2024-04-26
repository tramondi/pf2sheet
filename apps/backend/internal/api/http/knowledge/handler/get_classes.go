package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	common_dto "github.com/alionapermes/pf2sheet/internal/api/http/common/dto"
	"github.com/alionapermes/pf2sheet/internal/api/http/knowledge/dto"
	"github.com/alionapermes/pf2sheet/internal/app/resource"
)

func GetClasses(container resource.Container) echo.HandlerFunc {
	type data struct {
		Classes []dto.Class `json:"classes"`
	}

	getAllClasses := container.Usecases.GetAllClasses

	return func(ctx echo.Context) error {
		classes, err := getAllClasses.Execute(context.Background())
		if err != nil {
			return ctx.String(http.StatusInternalServerError, "internal error")
		}

		DTOs := make([]dto.Class, 0, len(classes))
		for _, class := range classes {
			DTOs = append(DTOs, dto.Class{
				ID:    int64(class.ID),
				Title: class.Title,
			})
		}

		response := &common_dto.Response{
			Data: &data{Classes: DTOs},
		}

		return ctx.JSON(http.StatusOK, response)
	}
}
