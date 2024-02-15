package services

import (
	"context"
	"go-junior-test/stock/internal/models"
)

type stocker interface {
	GetStatusStockAvailability(ctx context.Context, stockId uint32) (models.Status, error)
	ReserveList(ctx context.Context, items []*models.ItemReserve) error
	ReserveCancelList(ctx context.Context, items []*models.ItemReserveCancel) error
	GetItemsByStock(ctx context.Context, stockId uint32) (item []*models.ItemStock, err error)
	ArrivalList(ctx context.Context, items []*models.ItemArrival) error
	GetAvailabilityBySKUAndStockID(ctx context.Context, sku, stockId uint32) (uint32, uint32, error)
}

type service struct {
	stockClient stocker
}

func NewService(stockClient stocker) *service {
	return &service{
		stockClient,
	}
}
