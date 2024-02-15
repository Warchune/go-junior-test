// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package stock

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	Arrival(ctx context.Context, sku int32, stockID int32, available pgtype.Int4) error
	GetItemsByStock(ctx context.Context, stockID int32) error
	GetStatusStockAvailability(ctx context.Context, id int32) error
	Reserve(ctx context.Context, sku int32, available pgtype.Int4, stockID int32) error
	ReserveCancel(ctx context.Context, sku int32, available pgtype.Int4, stockID int32) error
	UpdateItem(ctx context.Context, sku int32, availableAll pgtype.Int4, reservedAll pgtype.Int4) error
}

var _ Querier = (*Queries)(nil)
