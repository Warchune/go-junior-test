-- +goose Up

insert into statuses(id, status)
values  (1, 'available'),
        (2, 'not available');

insert into stocks(id, name, status_id)
values (1, 'mosk01', 1);

insert into items(name, size, sku, available_all, reserved_all)
values ('футболка', 'M', 0010, 60, 80),
       ('шорты', 'XXL', 0831, 10, 2),
       ('джинсы', 'L', 9014, 200, 50),
       ('шапка', '52', 7770, 41, 8),
       ('куртка', 'S', 7100, 8, 1);

insert into item_stock(sku, stock_id, available, reserved)
values (0010, 1, 60, 20),
       (0831, 1, 10, 2),
       (9014, 1, 200, 50),
       (7770, 1, 41, 8),
       (7100, 1, 8, 1);

-- +goose Down

truncate stock_items;
truncate items;
truncate stocks;
truncate statuses;