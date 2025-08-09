-- +goose Up
-- +goose StatementBegin
alter table products add column deleted_at timestamp;

alter table stocks_locations add column slug varchar(255);

create index idx_stocks_location_slug on stocks_locations (id, slug);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table products drop column if exists deleted_at;

DROP INDEX IF EXISTS idx_stocks_location_slug;

ALTER TABLE stocks_locations DROP COLUMN slug;
-- +goose StatementEnd
