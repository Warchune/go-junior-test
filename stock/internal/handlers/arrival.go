package handlers

import (
	"context"
	"encoding/json"
	"go-junior-test/stock/internal/models"
	"net/http"
	"time"
)

type ArrivalRequest struct {
	Items []struct {
		Name    string `json:"Name"`
		Size    string `json:"Size"`
		SKU     uint32 `json:"SKU,omitempty"`
		Count   uint32 `json:"Count,omitempty"`
		StockId uint32 `json:"StockId,omitempty"`
	}
}

func (r ArrivalRequest) Validate() error {
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

func (c *controller) Arrival(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 500*time.Millisecond)
	defer cancel()

	req := &ArrivalRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		GetErrorResponse(w, c.name, err, http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		GetErrorResponse(w, c.name, err, http.StatusBadRequest)
		return
	}

	itemsArrival := make([]*models.ItemArrival, 0)
	for _, item := range req.Items {
		itemsArrival = append(itemsArrival, &models.ItemArrival{
			Name:    item.Name,
			Size:    item.Size,
			SKU:     item.SKU,
			Count:   item.Count,
			StockId: item.StockId,
		})
	}

	if err := c.Service.Arrival(ctx, itemsArrival); err != nil {
		GetErrorResponse(w, c.name, err, http.StatusInternalServerError)
	}
}
