package http

import (
	"go-junior-test/stock/internal/clients"
	"go-junior-test/stock/internal/handlers"
	"go-junior-test/stock/internal/services"
	"log"
	"net/http"
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
	Client := clients.NewClient()

	Controller := handlers.NewController(services.NewService(Client))

	http.HandleFunc("/stock/reserve", Controller.Reserve)
	http.HandleFunc("/stock/reserve_cancel", Controller.ReserveCancel)
	http.HandleFunc("/stock/info", Controller.Info)
	http.HandleFunc("stock/arrival", Controller.Arrival)

	log.Printf("http слушает по адресу %s", a.cfg.Address)

	return http.ListenAndServe(a.cfg.Address, nil)
}
