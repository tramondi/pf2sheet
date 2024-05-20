package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/AlekSi/pointer"
	"github.com/labstack/echo/v4"

	"github.com/alionapermes/pf2sheet/internal/api/http/common"
	common_dto "github.com/alionapermes/pf2sheet/internal/api/http/common/dto"
	"github.com/alionapermes/pf2sheet/internal/app/resource"
	"github.com/alionapermes/pf2sheet/internal/domain/contract"
)

func Signup(container resource.Container) echo.HandlerFunc {
	type credentials struct {
		Login       string  `json:"login"`
		Password    string  `json:"password"`
		DisplayName *string `json:"display_name,omitempty"`
	}

	signup := container.Usecases.Signup

	return func(ctx echo.Context) error {
		var creds credentials
		if err := ctx.Bind(&creds); err != nil {
			ctx.Logger().Errorf("failed to bind creds: %s", err)
			return ctx.NoContent(http.StatusBadRequest)
		}

		login := common.ClearText(creds.Login)
		displayName := common.ClearTextPtr(creds.DisplayName)

		session, err := signup.Execute(
			context.Background(), login, creds.Password, displayName)
		if err != nil {
			ctx.Logger().Errorf("failed to execute signup usecase: %s", err)

			if errors.Is(err, contract.ErrAlreadyExists) {
				msg := fmt.Sprintf("player with login %s already exists", creds.Login)
				response := &common_dto.Response{
					Error: pointer.ToString(msg),
				}

				return ctx.JSON(http.StatusConflict, response)
			}

			return ctx.NoContent(http.StatusInternalServerError)
		}

		return setSessionCookie(ctx, session).
			NoContent(http.StatusOK)
	}
}
