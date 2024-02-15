package services

import (
	"context"
	"errors"
	"go-junior-test/stock/internal/models"
)

func (s *service) ReserveCancel(ctx context.Context, items []*models.ItemReserveCancel) error {
	for _, item := range items {
		status, err := s.stockClient.GetStatusStockAvailability(ctx, item.StockId)
		if err != nil {

			return err
		}
		if status != models.StatusAvailable {
			return errors.New("stock is not available")
		}

		_, reserved, err := s.stockClient.GetAvailabilityBySKUAndStockID(ctx, item.SKU, item.StockId)
		if err != nil {
			return err
		}
		if reserved <= item.Count {
			return errors.New("not enough items available to reserve cancel")
		}
	}

	return s.stockClient.ReserveCancelList(ctx, items)
}
