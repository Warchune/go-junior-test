package handlers

import (
	"context"
	"encoding/json"
	"go-junior-test/stock/internal/models"
	"net/http"
	"time"
)

type ReserveCancelRequest struct {
	Items []struct {
		SKU     uint32 `json:"SKU,omitempty"`
		Count   uint32 `json:"Count,omitempty"`
		StockId uint32 `json:"StockId,omitempty"`
	}
}

func (r ReserveCancelRequest) Validate() error {
	for _, item := range r.Items {
		if item.SKU <= 0 {
			return ErrIncorrectSKU
		}
		if item.Count <= 0 {
			return ErrIncorrectQuantity
		}
		if item.StockId <= 0 {
			return ErrIncorrectStockId
		}
	}
	return nil
}

func (c *controller) ReserveCancel(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 500*time.Millisecond)
	defer cancel()

	req := &ReserveCancelRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		GetErrorResponse(w, c.name, err, http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		GetErrorResponse(w, c.name, err, http.StatusBadRequest)
		return
	}

	itemsReserveCancel := make([]*models.ItemReserveCancel, len(req.Items))
	for _, item := range req.Items {
		itemsReserveCancel = append(itemsReserveCancel, &models.ItemReserveCancel{
			SKU:     item.SKU,
			Count:   item.Count,
			StockId: item.StockId,
		})
	}

	if err := c.Service.ReserveCancel(ctx, itemsReserveCancel); err != nil {
		GetErrorResponse(w, c.name, err, http.StatusInternalServerError)
	}
}
