-- +goose Up
-- +goose StatementBegin
create table outboxes
(
    id uuid primary key not null,
    exchange varchar(255) not null,
    event_type varchar(255) not null,
    payload jsonb not null,
    created_at timestamp default now(),
    processed bool not null default false,
    processed_at timestamp
);

create table processed_messages
(
    message_id uuid primary key not null
);

alter table products add column slug varchar(255);

create index idx_vendor_product_slug on products (vendor_id, slug);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP INDEX IF EXISTS idx_vendor_product_slug;

ALTER TABLE products DROP COLUMN slug;

DROP TABLE IF EXISTS outboxes;
DROP TABLE IF EXISTS processed_messages;
-- +goose StatementEnd
