-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS box_transactions (
    id CHAR(36) NOT NULL PRIMARY KEY,
    type VARCHAR(10) NOT NULL,
    quantity DECIMAL(22,6) NOT NULL,
    box_id CHAR(36) NOT NULL,
    item_id CHAR(36) NOT NULL,
    item_sku VARCHAR(100) NOT NULL,
    item_name VARCHAR(100) NOT NULL,
    item_unit VARCHAR(10) NOT NULL,
    happened_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT box_transactions_box_id_fk FOREIGN KEY (box_id) REFERENCES boxes(id),
    CONSTRAINT box_transactions_item_id_fk FOREIGN KEY (item_id) REFERENCES items(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE box_transactions;
-- +goose StatementEnd
