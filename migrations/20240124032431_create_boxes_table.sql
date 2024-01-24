-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS boxes (
    id CHAR(36) NOT NULL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(255) NULL,
    room_id CHAR(36) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT boxes_room_id_fk FOREIGN KEY (room_id) REFERENCES rooms(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE boxes;
-- +goose StatementEnd
