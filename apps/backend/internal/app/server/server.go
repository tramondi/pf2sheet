package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	echo_middleware "github.com/labstack/echo/v4/middleware"
	"github.com/lmittmann/tint"

	"github.com/alionapermes/pf2sheet/internal/app/resource"
	"github.com/alionapermes/pf2sheet/internal/infra/sqlc-pg/dao"
)

type Server struct {
	logger *slog.Logger
	config Config

	e *echo.Echo
}

func (self *Server) Start() {
	self.e = echo.New()
	self.e.Use(echo_middleware.Logger())
	self.e.Use(echo_middleware.CORSWithConfig(echo_middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodOptions,
		},
		AllowHeaders: []string{
			echo.HeaderAccept,
			echo.HeaderContentType,
			echo.HeaderOrigin,
		},
	}))

	self.logger = slog.New(tint.NewHandler(os.Stderr, nil))

	// #####
	const dsn = "postgres://postgres:postgres@db:5432/postgres?sslmode=disable&client_encoding=UTF8"

	pgxConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		panic(err)
	}

	pgxConfig.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	db, err := pgxpool.NewWithConfig(context.Background(), pgxConfig)
	if err != nil {
		panic(err)
	}

	container, err := resource.InitContainer(self.logger, dao.New(db))
	if err != nil {
		panic(err)
	}
	// #####

	self.initRoutes(container)

	self.e.Start(":8080")
}
