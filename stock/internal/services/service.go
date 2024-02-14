package services

import (
	"context"
	"go-junior-test/stock/internal/models"
)

type stocker interface {
	GetStatusAvailability(ctx context.Context, stockId uint32) (models.Status, error)
	Reserve(ctx context.Context, sku uint32, count uint32, stockId uint32) error
	ReserveCancel(ctx context.Context, sku uint32, count uint32, stockId uint32) error
	GetBySku(ctx context.Context, sku uint32) (item *models.ItemStock, err error)
	Arrival(ctx context.Context, sku uint32, count uint32, stockId uint32) error
}

type service struct {
	stockClient stocker
}

func NewService(stockClient stocker) *service {
	return &service{
		stockClient,
	}
}
