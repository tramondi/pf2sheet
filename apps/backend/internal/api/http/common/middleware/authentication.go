package middleware

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/alionapermes/pf2sheet/internal/app/resource"
)

func Authentication(container resource.Container) echo.MiddlewareFunc {
	sessionsRepo := container.Repositories.Sessions

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			cookie, err := ctx.Cookie("session_token")
			if err != nil {
				return ctx.NoContent(http.StatusUnauthorized)
			}

			token := cookie.Value
			if token == "" {
				return ctx.NoContent(http.StatusUnauthorized)
			}

			session, err := sessionsRepo.FindByToken(context.Background(), token)
			if err != nil {
				return ctx.NoContent(http.StatusUnauthorized)
			}

			ctx.Set("player_id", session.PlayerID)

			return next(ctx)
		}
	}
}
