package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
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
	res, err := c.q.GetStatusStockAvailability(ctx, int32(stockId))
	if err != nil {
		return "", nil
	}
	return models.Status(res), nil
}

func (c *client) ArrivalList(ctx context.Context, items []*models.ItemArrival) error {
	tx, err := c.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()
	q := c.q.WithTx(tx)
	for _, item := range items {
		arg := &pg.ArrivalParams{
			Name:      pgtype.Text{String: item.Name},
			Size:      pgtype.Text{String: item.Size},
			Sku:       int32(item.SKU),
			Available: pgtype.Int4{Int32: int32(item.Count), Valid: true},
			StockID:   int32(item.StockId),
		}
		if err := q.Arrival(ctx, arg); err != nil {
			return err
		}
	}
	return tx.Commit(ctx)
}

func (c *client) ReserveList(ctx context.Context, items []*models.ItemReserve) error {
	tx, err := c.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()
	q := c.q.WithTx(tx)
	for _, item := range items {
		if err := q.Reserve(ctx, int32(item.SKU), pgtype.Int4{Int32: int32(item.Count), Valid: true}, int32(item.StockId)); err != nil {
			return err
		}
		if err := q.UpdateItem(ctx, int32(item.SKU), pgtype.Int4{Int32: int32(0 - item.Count)}, pgtype.Int4{Int32: int32(item.Count)}); err != nil {
			return err
		}
	}
	return tx.Commit(ctx)
}

func (c *client) ReserveCancelList(ctx context.Context, items []*models.ItemReserveCancel) error {
	tx, err := c.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()
	q := c.q.WithTx(tx)
	for _, item := range items {
		if err := q.ReserveCancel(ctx, int32(item.SKU), pgtype.Int4{Int32: int32(item.Count), Valid: true}, int32(item.StockId)); err != nil {
			return err
		}
		if err := q.UpdateItem(ctx, int32(item.SKU), pgtype.Int4{Int32: int32(item.Count)}, pgtype.Int4{Int32: int32(0 - item.Count)}); err != nil {
			return err
		}
	}
	return tx.Commit(ctx)
}

func (c *client) GetItemsByStock(ctx context.Context, stockId uint32) (item []*models.ItemStock, err error) {
	res, err := c.q.GetItemsByStock(ctx, int32(stockId))
	if err != nil {
		return nil, err
	}
	items := make([]*models.ItemStock, 0)
	for _, item := range res {
		items = append(items, &models.ItemStock{
			SKU:       uint32(item.Sku),
			Available: uint32(item.Available.Int32),
			Reserved:  uint32(item.Reserved.Int32),
			StockId:   stockId,
		})
	}
	return items, nil
}

func (c *client) GetAvailabilityBySKUAndStockID(ctx context.Context, sku, stockId uint32) (uint32, uint32, error) {
	res, err := c.q.GetAvailabilityBySKUAndStockID(ctx, int32(sku), int32(stockId))
	if err != nil {
		return 0, 0, err
	}
	return uint32(res.Available.Int32), uint32(res.Reserved.Int32), nil
}
