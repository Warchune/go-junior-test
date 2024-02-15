-- name: GetStatusStockAvailability :exec
select s.status
from stocks st
    join statuses s on st.status_id = s.id
where st.id = $1;

-- name: Arrival :exec
insert into item_stock (sku, stock_id, available, reserved)
values ($1, $2, $3, 0)
on conflict (sku, stock_id) do update
set available = item_stock.available + EXCLUDED.available;

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

-- name: GetItemsByStock :exec
select i.name, i.size, ist.available, ist.reserved
from item_stock ist
    join items i on ist.sku = i.sku
where ist.stock_id = $1;