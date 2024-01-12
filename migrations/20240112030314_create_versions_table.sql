-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS versions (
    id CHAR(36) NOT NULL PRIMARY KEY,
    `value` VARCHAR(11)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE versions;
-- +goose StatementEnd
