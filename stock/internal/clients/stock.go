package clients

import (
	"context"
	"go-junior-test/stock/internal/models"
)

type client struct {
}

func NewClient() *client {
	return &client{}
}

func (c *client) GetStatusAvailability(ctx context.Context, stockId uint32) (models.Status, error) {
	return "", nil
}

func (c *client) Arrival(ctx context.Context, sku uint32, count uint32, stockId uint32) error {
	return nil
}

func (c *client) Reserve(ctx context.Context, sku uint32, count uint32, stockId uint32) error {
	return nil
}

func (c *client) ReserveCancel(ctx context.Context, sku uint32, count uint32, stockId uint32) error {
	return nil
}

func (c *client) GetBySku(ctx context.Context, sku uint32) (item *models.ItemStock, err error) {
	return nil, nil
}
