package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

const sessionCookieName = "session_token"

func setSessionCookie(ctx echo.Context, session entity.Session) echo.Context {
	ctx.SetCookie(&http.Cookie{
		Name:     sessionCookieName,
		Value:    session.Token,
		Path:     "/",
		HttpOnly: true,
	})

	return ctx
}
