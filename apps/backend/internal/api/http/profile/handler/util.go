package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

const sessionCookieName = "session_token"

func unsetSessionCookie(ctx echo.Context) echo.Context {
	ctx.SetCookie(&http.Cookie{
		Name:     sessionCookieName,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Time{},
	})

	return ctx
}
