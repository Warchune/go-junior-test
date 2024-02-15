package services

import (
	"context"
	"errors"
	"go-junior-test/stock/internal/models"
)

func (s *service) Arrival(ctx context.Context, items []*models.ItemArrival) error {
	for _, item := range items {
		status, err := s.stockClient.GetStatusStockAvailability(ctx, item.StockId)
		if err != nil {
			return err
		}
		if status != models.StatusAvailable {
			return errors.New("stock is not available")
		}
	}

	return s.stockClient.ArrivalList(ctx, items)
}
