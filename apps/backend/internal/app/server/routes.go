package server

import (
	"net/http"

	"github.com/labstack/echo/v4"

	auth_handler "github.com/alionapermes/pf2sheet/internal/api/http/auth/handler"
	auth_middleware "github.com/alionapermes/pf2sheet/internal/api/http/auth/middleware"
	common_middleware "github.com/alionapermes/pf2sheet/internal/api/http/common/middleware"
	knowledge_handler "github.com/alionapermes/pf2sheet/internal/api/http/knowledge/handler"
	"github.com/alionapermes/pf2sheet/internal/app/resource"
)

func (self *Server) initRoutes(container resource.Container) {
	self.e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	authAPI := self.e.Group("/auth", auth_middleware.Authentication(container))

	authAPI.POST("/signin", auth_handler.Signin(container))
	authAPI.POST("/signup", auth_handler.Signup(container))

	groupAPI := self.e.Group("/api", common_middleware.Authentication(container))

	groupKnowledge := groupAPI.Group("/knowledge")
	groupKnowledge.GET("/ancestries", knowledge_handler.GetAncestries(container))
	groupKnowledge.GET("/classes", knowledge_handler.GetClasses(container))
}
