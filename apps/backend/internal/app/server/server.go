package server

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/alionapermes/pf2sheet/internal/adapter/pg"
	"github.com/alionapermes/pf2sheet/internal/api/http/knowledge/handler"
	"github.com/alionapermes/pf2sheet/internal/app/resource"
	"github.com/alionapermes/pf2sheet/internal/app/usecase"
	"github.com/alionapermes/pf2sheet/internal/infra/sqlc-pg/dao"
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

	self.initRoutes()

	self.e.Start(":8080")
}

func (self *Server) initRoutes() {
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

	knowledgeRepo := pg.NewKnowledgeRepo(nil, dao.New(db))

	usecases := resource.Usecases{
		GetAllAncestries: usecase.NewGetAllAncestriesUsecase(nil, knowledgeRepo),
	}

	self.e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	groupAPI := self.e.Group("/api")

	groupKnowledge := groupAPI.Group("/knowledge")
	groupKnowledge.GET("/ancestries", handler.GetAncestries(&usecases))
}
