package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/AlekSi/pointer"
	"github.com/labstack/echo/v4"

	common_dto "github.com/alionapermes/pf2sheet/internal/api/http/common/dto"
	"github.com/alionapermes/pf2sheet/internal/app/resource"
	"github.com/alionapermes/pf2sheet/internal/domain/contract"
)

func Signup(container resource.Container) echo.HandlerFunc {
	type credentials struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	signup := container.Usecases.Signup

	return func(ctx echo.Context) error {
		var creds credentials
		if err := ctx.Bind(&creds); err != nil {
			return ctx.NoContent(http.StatusBadRequest)
		}

		session, err := signup.Execute(
			context.Background(), creds.Login, creds.Password)
		if err != nil {
			if errors.Is(err, contract.ErrAlreadyExists) {
				msg := fmt.Sprintf("player with login %s already exists", creds.Login)
				response := &common_dto.Response[any]{
					Error: pointer.ToString(msg),
				}

				return ctx.JSON(http.StatusConflict, response)
			}

			return ctx.NoContent(http.StatusInternalServerError)
		}

		ctx.SetCookie(&http.Cookie{
			Name:  "session_token",
			Value: session.Token,
		})

		return ctx.NoContent(http.StatusOK)
	}
}
