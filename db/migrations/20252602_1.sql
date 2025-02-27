-- noinspection SqlDialectInspectionForFile

-- +goose Up
-- +goose StatementBegin
-- noinspection SqlDialectInspection

CREATE TABLE IF NOT EXISTS images (
    name VARCHAR(255) PRIMARY KEY,
    data BYTEA NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT NULL
    );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS images;

-- +goose StatementEnd