package services

import (
	"context"
	"errors"
	"go-junior-test/stock/internal/models"
	"strconv"
)

func (s *service) ReserveCancel(ctx context.Context, items []*models.ItemReserveCancel) error {
	for _, item := range items {
		status, err := s.stockClient.GetStatusStockAvailability(ctx, item.StockId)
		if err != nil {
			return err
		}
		if status == models.StatusAvailable {
			continue
		} else {
			err := errors.New(strconv.Itoa(int(item.StockId)) + string(models.StatusNotAvailable))
			return err
		}
	}

	return s.stockClient.ReserveCancelList(ctx, items)
}
