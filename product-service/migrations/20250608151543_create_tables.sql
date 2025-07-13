-- +goose Up
-- +goose StatementBegin

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table stocks_locations
(
    id      uuid         not null
        primary key,
    city    varchar(255) not null,
    address varchar(255) not null
);

create table products
(
    id          uuid                not null
        primary key,
    vendor_id   uuid                not null,
    name        varchar(255)        not null,
    description text                not null,
    price       numeric(12, 2)      not null,
    category    varchar(255)        not null,
    quantity    integer   default 0 not null,
    created_at  timestamp default now(),
    updated_at  timestamp default now()
);

create table tags
(
    id       uuid not null primary key,
    tag_name text not null
);

create table products_tags
(
    product_id uuid not null,
    tag_id     uuid not null,
    primary key (product_id, tag_id)
);

create table products_images
(
    id         uuid not null primary key,
    image_url  text not null,
    product_id uuid not null
);

create table stocks
(
    date_supplied date not null,
    id            uuid not null
        primary key,
    created_at    timestamp default now(),
    updated_at    timestamp default now(),
    location_id   uuid not null
);

create table stocks_products
(
    product_id uuid not null,
    unit_cost  numeric(12, 2),
    quantity   integer not null,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    stock_id   uuid not null,
    primary key (product_id, stock_id)
);

create index idx_products_category on products (category);

create index idx_products_name on products (name);

create index idx_products_price on products (price);

create index idx_products_quantity on products (quantity);

create index idx_products_vendor_id on products (vendor_id);


alter table products_images
    add constraint products_images_product_id_fkey
        foreign key (product_id) references products (id) on delete cascade;


alter table products_tags
    add constraint products_tags_product_id_fkey
        foreign key (product_id) references products (id) on delete cascade;


alter table products_tags
    add constraint products_tags_tag_id_fkey
        foreign key (tag_id) references tags (id) on delete cascade;

alter table stocks_products
    add constraint stocks_products_product_id_fkey
        foreign key (product_id) references products(id) on delete cascade;

alter table stocks_products
    add constraint stocks_products_stock_id_fkey
        foreign key (stock_id) references stocks(id) on delete cascade;

alter table stocks
    add column vendor_id uuid;

create index idx_stocks_vendor_id on stocks (vendor_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table products_images
drop constraint if exists products_images_product_id_fkey;

alter table products_tags
drop constraint if exists products_tags_product_id_fkey;

alter table stocks_products
drop constraint if exists stocks_products_product_id_fkey;

alter table products_tags
drop constraint if exists products_tags_tag_id_fkey;

alter table stocks_products
drop constraint if exists stocks_products_stock_id_fkey;

DROP INDEX IF EXISTS idx_products_price;
DROP INDEX IF EXISTS idx_products_category;
DROP INDEX IF EXISTS idx_products_name;
DROP INDEX IF EXISTS idx_products_quantity;
DROP INDEX IF EXISTS idx_products_vendor_id;
DROP INDEX IF EXISTS idx_stocks_vendor_id;

DROP TABLE IF EXISTS stocks_products;
DROP TABLE IF EXISTS stocks;
DROP TABLE IF EXISTS products_tags;
DROP TABLE IF EXISTS tags;
DROP TABLE IF EXISTS products_images;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS stocks_locations;
-- +goose StatementEnd
