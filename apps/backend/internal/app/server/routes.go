package server

import (
	"net/http"

	"github.com/labstack/echo/v4"

	auth_handler "github.com/alionapermes/pf2sheet/internal/api/http/auth/handler"
	auth_middleware "github.com/alionapermes/pf2sheet/internal/api/http/auth/middleware"
	common_middleware "github.com/alionapermes/pf2sheet/internal/api/http/common/middleware"
	knowledge_handler "github.com/alionapermes/pf2sheet/internal/api/http/knowledge/handler"
	profile_handler "github.com/alionapermes/pf2sheet/internal/api/http/profile/handler"
	sheets_handler "github.com/alionapermes/pf2sheet/internal/api/http/sheets/handler"
	"github.com/alionapermes/pf2sheet/internal/app/resource"
)

func (self *Server) initRoutes(container resource.Container) {
	commonAuthMiddleware := common_middleware.Authentication(container)

	self.e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	groupAuth := self.e.Group("/auth", auth_middleware.Authentication(container))
	{
		groupAuth.POST("/signin", auth_handler.Signin(container))
		groupAuth.POST("/signup", auth_handler.Signup(container))
	}

	groupProfile := self.e.Group("/profile", commonAuthMiddleware)
	{
		groupProfile.GET("/", profile_handler.GetProfile(container))
		groupProfile.POST("/signout", profile_handler.Signout(container))
		groupProfile.PUT("/", profile_handler.UpdateProfile(container))
		groupProfile.DELETE("/", profile_handler.DeleteProfile(container))
	}

	groupAPI := self.e.Group("/api", commonAuthMiddleware)
	{
		groupKnowledge := groupAPI.Group("/knowledge")
		{
			groupKnowledge.GET(
				"/ancestries", knowledge_handler.GetAncestries(container))
			groupKnowledge.GET("/classes", knowledge_handler.GetClasses(container))
		}

		groupSheets := groupAPI.Group("/sheets")
		{
			groupSheets.POST("/", sheets_handler.CreateSheet(container))
			groupSheets.GET("/", sheets_handler.GetAllSheets(container))
			groupSheets.PUT("/:id", sheets_handler.UpdateSheet(container))
			groupSheets.DELETE("/:id", sheets_handler.DeleteSheet(container))
		}
	}
}
