package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/alionapermes/pf2sheet/internal/api/http/auth/dto"
	"github.com/alionapermes/pf2sheet/internal/app/resource"
	"github.com/alionapermes/pf2sheet/internal/domain/contract"
)

func Authentication(container resource.Container) echo.MiddlewareFunc {
	sessionsRepo := container.Repositories.Sessions

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			cookie, err := ctx.Cookie("session_token")
			if errors.Is(err, http.ErrNoCookie) || cookie.Value == "" {
				return next(ctx)
			}

			_, err = sessionsRepo.FindByToken(context.Background(), cookie.Value)
			if errors.Is(err, contract.ErrNotFound) {
				return next(ctx)
			}

			return ctx.String(
				http.StatusConflict, dto.ErrSessionAlreadyExists.Error())
		}
	}
}
