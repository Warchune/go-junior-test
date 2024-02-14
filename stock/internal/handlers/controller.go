package handlers

import (
	"context"
	"go-junior-test/stock/internal/models"
)

type stocker interface {
	Reserve(ctx context.Context, items []*models.ItemReserve) error
	ReserveCancel(ctx context.Context, items []*models.ItemReserveCancel) error
	Info(ctx context.Context, stockId uint32) (items []*models.ItemStock, err error)
	Arrival(ctx context.Context, items []*models.Item) error
}

type controller struct {
	name    string
	Service stocker
}

func NewController(stockService stocker) *controller {
	return &controller{
		"controller",
		stockService,
	}
}
