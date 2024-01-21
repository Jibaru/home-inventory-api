-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS assets (
    id CHAR(36) NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    extension VARCHAR(10) NOT NULL,
    size BIGINT NOT NULL,
    file_id CHAR(36) NOT NULL,
    entity_id CHAR(36) NOT NULL,
    entity_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    INDEX assets_entity_id_entity_name_idx (entity_id, entity_name)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE assets;
-- +goose StatementEnd
