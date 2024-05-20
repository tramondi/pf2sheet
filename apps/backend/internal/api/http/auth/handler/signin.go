package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/alionapermes/pf2sheet/internal/app/resource"
)

func Signin(container resource.Container) echo.HandlerFunc {
	type credentials struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	signin := container.Usecases.Signin

	return func(ctx echo.Context) error {
		var creds credentials
		if err := ctx.Bind(&creds); err != nil {
			ctx.Logger().Errorf("failed to bind creds: %s", err)
			return ctx.NoContent(http.StatusBadRequest)
		}

		session, err := signin.Execute(
			context.Background(),
			creds.Login,
			creds.Password,
		)
		if err != nil {
			ctx.Logger().Errorf("failed to execute signin usecase: %s", err)
			return ctx.NoContent(http.StatusUnauthorized)
		}

		return setSessionCookie(ctx, session).
			NoContent(http.StatusOK)
	}
}
