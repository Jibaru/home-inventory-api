-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS box_items (
    id CHAR(36) NOT NULL PRIMARY KEY,
    quantity DECIMAL(22,6) NOT NULL,
    box_id CHAR(36) NOT NULL,
    item_id CHAR(36) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT box_items_box_id_fk FOREIGN KEY (box_id) REFERENCES boxes(id),
    CONSTRAINT box_items_item_id_fk FOREIGN KEY (item_id) REFERENCES items(id),
    CONSTRAINT box_items_box_id_item_id_unq UNIQUE (box_id, item_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE box_items;
-- +goose StatementEnd
