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
				ctx.Logger().Error("failed to get session token from ctx")
				return ctx.NoContent(http.StatusUnauthorized)
			}

			token := cookie.Value
			if token == "" {
				ctx.Logger().Error("session token is empty")
				return ctx.NoContent(http.StatusUnauthorized)
			}

			session, err := sessionsRepo.GetByToken(context.Background(), token)
			if err != nil {
				ctx.Logger().Errorf("failed to get session by token: %s", err)
				return ctx.NoContent(http.StatusUnauthorized)
			}

			ctx.Set("player_id", session.PlayerID)
			ctx.Set("session_token", session.Token)

			return next(ctx)
		}
	}
}
