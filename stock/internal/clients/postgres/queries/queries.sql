-- name: GetStatusStockAvailability :one
select s.status
from stocks st
    join statuses s on st.status_id = s.id
where st.id = $1;

-- name: Arrival :exec
with inserted_item as (
insert into items (name, size, sku, available_all, reserved_all)
values ($1, $2, $3, $4, 0)
on conflict (sku) do update
set available_all = items.available_all + excluded.available_all
returning sku
)
insert into item_stock (sku, stock_id, available, reserved)
select sku, $5, $4, 0
from inserted_item
on conflict (sku, stock_id) do update
set available = item_stock.available + excluded.available,
reserved = item_stock.reserved;

-- name: Reserve :exec
update item_stock
set available = available - $2,
    reserved = reserved + $2
where sku = $1 and stock_id = $3;

-- name: ReserveCancel :exec
update item_stock
set available = available + $2,
    reserved = reserved - $2
where sku = $1 and stock_id = $3;

-- name: UpdateItem :exec
update items
set available_all = available_all + $2,
    reserved_all = reserved_all + $3
where sku = $1;

-- name: GetItemsByStock :many
select i.sku, i.size, ist.available, ist.reserved
from item_stock ist
    join items i on ist.sku = i.sku
where ist.stock_id = $1;