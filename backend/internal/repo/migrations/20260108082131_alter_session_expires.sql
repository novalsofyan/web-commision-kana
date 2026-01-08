-- +goose Up
-- +goose StatementBegin
ALTER TABLE sessions 
ADD COLUMN expires_at TIMESTAMPTZ NOT NULL 
DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC' AT TIME ZONE 'Asia/Jakarta' + INTERVAL '30 days');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE sessions 
DROP COLUMN IF EXISTS expires_at;
-- +goose StatementEnd