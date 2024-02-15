package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	pg "go-junior-test/stock/internal/clients/postgres/internal/gen"
	"go-junior-test/stock/internal/models"
)

type client struct {
	pool *pgxpool.Pool
	q    *pg.Queries
}

func NewClient(pool *pgxpool.Pool) *client {
	return &client{
		pool: pool,
		q:    pg.New(pool),
	}
}
func (c *client) GetStatusStockAvailability(ctx context.Context, stockId uint32) (models.Status, error) {
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

func (c *client) GetItemsByStock(ctx context.Context, stockId uint32) (item []*models.ItemStock, err error) {
	return nil, nil
}
