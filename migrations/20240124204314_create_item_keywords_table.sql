-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS item_keywords (
    id CHAR(36) NOT NULL PRIMARY KEY,
    value VARCHAR(60) NOT NULL,
    item_id CHAR(36) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT item_keywords_item_id_fk FOREIGN KEY (item_id) REFERENCES items(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE item_keywords;
-- +goose StatementEnd
