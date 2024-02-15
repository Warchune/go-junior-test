-- +goose Up

create table statuses (
    id serial primary key,
    status text not null
);

create table stocks (
    id serial primary key,
    name varchar not null,
    status_id int4 not null,
    foreign key (status_id) references statuses(id)
);

create table items (
    name varchar,
    size varchar,
    sku integer not null primary key,
    available_all int4,
    reserved_all int4
);

create table item_stock (
    sku integer not null ,
    stock_id integer not null ,
    available int4,
    reserved int4,
    primary key (sku, stock_id),
    foreign key (sku) references items(sku),
    foreign key (stock_id) references stocks(id)
);

-- +goose Down

drop table if exists items;
drop table if exists item_stock;
drop table if exists statuses;
drop table if exists stocks;
