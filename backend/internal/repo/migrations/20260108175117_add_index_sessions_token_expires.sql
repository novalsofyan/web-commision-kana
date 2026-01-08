-- +goose Up
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_sessions_token;

CREATE INDEX idx_sessions_token_expires ON sessions(token, expires_at);

CREATE INDEX idx_products_user_id ON products(user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_products_user_id;
DROP INDEX IF EXISTS idx_sessions_token_expires;

CREATE INDEX idx_sessions_token ON sessions(token);
-- +goose StatementEnd