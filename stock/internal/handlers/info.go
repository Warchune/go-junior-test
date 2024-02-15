package handlers

import (
	"context"
	"encoding/json"
	"go-junior-test/stock/internal/models"
	"net/http"
	"time"
)

type InfoRequest struct {
	StockId uint32 `json:"StockId,omitempty"`
}

type InfoResponse struct {
	Items []*models.ItemStock
}

func (r InfoRequest) Validate() error {
	if r.StockId <= 0 {
		return ErrIncorrectStockId
	}

	return nil
}

func (c *controller) Info(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 500*time.Millisecond)
	defer cancel()

	req := &InfoRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		GetErrorResponse(w, c.name, err, http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		GetErrorResponse(w, c.name, err, http.StatusBadRequest)
		return
	}

	items, err := c.Service.Info(ctx, req.StockId)
	if err != nil {
		GetErrorResponse(w, c.name, err, http.StatusInternalServerError)
	}

	res := &InfoResponse{
		Items: items,
	}
	if err := WriteJSON(w, http.StatusOK, res); err != nil {
		GetErrorResponse(w, c.name, err, http.StatusInternalServerError)
		return
	}
}
