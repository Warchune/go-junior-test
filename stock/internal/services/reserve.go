package services

import (
	"context"
	"errors"
	"go-junior-test/stock/internal/models"
)

func (s *service) Reserve(ctx context.Context, items []*models.ItemReserve) error {
	for _, item := range items {
		status, err := s.stockClient.GetStatusStockAvailability(ctx, item.StockId)
		if err != nil {
			return err
		}

		if status != models.StatusAvailable {
			return errors.New("stock is not available")
		}

		available, _, err := s.stockClient.GetAvailabilityBySKUAndStockID(ctx, item.SKU, item.StockId)
		if err != nil {
			return err
		}
		if available <= item.Count {
			return errors.New("not enough items available to reserve")
		}
	}

	return s.stockClient.ReserveList(ctx, items)
}
