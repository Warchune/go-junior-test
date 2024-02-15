// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: queries.sql

package stock

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const arrival = `-- name: Arrival :exec
insert into item_stock (sku, stock_id, available, reserved)
values ($1, $2, $3, 0)
on conflict (sku, stock_id) do update
set available = item_stock.available + EXCLUDED.available
`

func (q *Queries) Arrival(ctx context.Context, sku int32, stockID int32, available pgtype.Int4) error {
	_, err := q.db.Exec(ctx, arrival, sku, stockID, available)
	return err
}

const getItemsByStock = `-- name: GetItemsByStock :exec
select i.name, i.size, ist.available, ist.reserved
from item_stock ist
    join items i on ist.sku = i.sku
where ist.stock_id = $1
`

func (q *Queries) GetItemsByStock(ctx context.Context, stockID int32) error {
	_, err := q.db.Exec(ctx, getItemsByStock, stockID)
	return err
}

const getStatusStockAvailability = `-- name: GetStatusStockAvailability :exec
select s.status
from stocks st
    join statuses s on st.status_id = s.id
where st.id = $1
`

func (q *Queries) GetStatusStockAvailability(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, getStatusStockAvailability, id)
	return err
}

const reserve = `-- name: Reserve :exec
update item_stock
set available = available - $2,
    reserved = reserved + $2
where sku = $1 and stock_id = $3
`

func (q *Queries) Reserve(ctx context.Context, sku int32, available pgtype.Int4, stockID int32) error {
	_, err := q.db.Exec(ctx, reserve, sku, available, stockID)
	return err
}

const reserveCancel = `-- name: ReserveCancel :exec
update item_stock
set available = available + $2,
    reserved = reserved - $2
where sku = $1 and stock_id = $3
`

func (q *Queries) ReserveCancel(ctx context.Context, sku int32, available pgtype.Int4, stockID int32) error {
	_, err := q.db.Exec(ctx, reserveCancel, sku, available, stockID)
	return err
}

const updateItem = `-- name: UpdateItem :exec
update items
set available_all = available_all + $2,
    reserved_all = reserved_all + $3
where sku = $1
`

func (q *Queries) UpdateItem(ctx context.Context, sku int32, availableAll pgtype.Int4, reservedAll pgtype.Int4) error {
	_, err := q.db.Exec(ctx, updateItem, sku, availableAll, reservedAll)
	return err
}
