-- +goose Up
-- +goose StatementBegin
-- Tabel Users
ALTER TABLE users ALTER COLUMN created_at TYPE TIMESTAMPTZ USING created_at AT TIME ZONE 'Asia/Jakarta';

-- Tabel Sessions
ALTER TABLE sessions ALTER COLUMN created_at TYPE TIMESTAMPTZ USING created_at AT TIME ZONE 'Asia/Jakarta';

-- Tabel Products
ALTER TABLE products ALTER COLUMN created_at TYPE TIMESTAMPTZ USING created_at AT TIME ZONE 'Asia/Jakarta';

-- Tabel Image Meta
ALTER TABLE image_meta ALTER COLUMN created_at TYPE TIMESTAMPTZ USING created_at AT TIME ZONE 'Asia/Jakarta';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE image_meta ALTER COLUMN created_at TYPE TIMESTAMP;
ALTER TABLE products ALTER COLUMN created_at TYPE TIMESTAMP;
ALTER TABLE sessions ALTER COLUMN created_at TYPE TIMESTAMP;
ALTER TABLE users ALTER COLUMN created_at TYPE TIMESTAMP;
-- +goose StatementEnd