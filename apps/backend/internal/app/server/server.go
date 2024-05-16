package server

import (
	"database/sql"
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

	e *echo.Echo
}

func (self *Server) Start() {
	self.e = echo.New()
	self.e.Use(middleware.Logger())
	self.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost"},
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

	// #####
	const dsn = "postgres://postgres:postgres@db:5432/postgres?sslmode=disable&client_encoding=UTF8"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	dialect := goqu.Dialect("postgres")
	goquDB := dialect.DB(db)

	container, err := resource.InitContainer(self.logger, goquDB)
	if err != nil {
		panic(err)
	}
	// #####

	self.initRoutes(container)

	self.e.Start(":8080")
}
