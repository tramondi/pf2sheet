package server

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/doug-martin/goqu/v9"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lmittmann/tint"

	_ "github.com/lib/pq"

	"github.com/alionapermes/pf2sheet/internal/app/resource"
)

type Server struct {
	logger *slog.Logger
	config Config

	router *echo.Echo
}

func (self *Server) Start(config Config) error {
	self.config = config

	self.router = echo.New()
	self.router.Use(middleware.Logger())
	self.router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: self.config.AllowedOrigins,
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowCredentials: true,
		AllowHeaders: []string{
			echo.HeaderAccept,
			echo.HeaderContentType,
			echo.HeaderOrigin,
		},
	}))

	self.logger = slog.New(tint.NewHandler(os.Stderr, nil))

	db, err := sql.Open("postgres", self.config.DSN)
	if err != nil {
		panic(err)
	}

	dialect := goqu.Dialect("postgres")
	goquDB := dialect.DB(db)

	container, err := resource.InitContainer(self.logger, goquDB)
	if err != nil {
		panic(err)
	}

	self.initRoutes(container)

	return self.router.Start(fmt.Sprintf(":%d", self.config.Port))
}
