-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS image_meta(
    id SERIAL PRIMARY KEY NOT NULL,
    nama_image VARCHAR(255) NOT NULL,
    url_image VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    product_id INTEGER NOT NULL,

    CONSTRAINT fk_product_image
        FOREIGN KEY(product_id)
        REFERENCES products(id)
        ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS image_meta;
-- +goose StatementEnd
