package http

import (
	"context"
	"go-junior-test/stock/internal/clients/postgres"
	"go-junior-test/stock/internal/handlers"
	"go-junior-test/stock/internal/services"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	cfg Config
}

func NewApp(cfg Config) *App {
	return &App{
		cfg: cfg,
	}
}

func (a *App) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pgPool, err := newPgPool(ctx, a.cfg.DatabaseURL)
	if err != nil {
		return err
	}

	Client := postgres.NewClient(pgPool)

	Controller := handlers.NewController(services.NewService(Client))

	http.HandleFunc("/stock/reserve", Controller.Reserve)
	http.HandleFunc("/stock/reserve_cancel", Controller.ReserveCancel)
	http.HandleFunc("/stock/info", Controller.Info)
	http.HandleFunc("/stock/arrival", Controller.Arrival)

	log.Printf("http слушает по адресу %s", a.cfg.Address)

	return http.ListenAndServe(a.cfg.Address, nil)
}

func newPgPool(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}
	cfg.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol //https://github.com/jackc/pgx/issues/1561
	return pgxpool.NewWithConfig(ctx, cfg)
}
