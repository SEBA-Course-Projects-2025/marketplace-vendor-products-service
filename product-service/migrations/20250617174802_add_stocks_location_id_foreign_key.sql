-- +goose Up
-- +goose StatementBegin
alter table stocks
    add constraint stocks_location_id_fkey foreign key (location_id)
        references stocks_locations(id) on delete cascade;

create index idx_products_id_vendor on products (id, vendor_id);

create index idx_stocks_id_vendor on stocks(id, vendor_id);
create index idx_stocks_products_stock_id on stocks_products(stock_id);
create index idx_stocks_products_stock_product on stocks_products(stock_id, product_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table stocks
    drop constraint stocks_location_id_fkey;

drop index if exists idx_products_id_vendor;
drop index if exists idx_stocks_id_vendor;
drop index if exists idx_stocks_products_stock_id;
drop index if exists idx_stocks_products_stock_product;
-- +goose StatementEnd
