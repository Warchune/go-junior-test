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
	}

	return s.stockClient.ReserveList(ctx, items)
}
