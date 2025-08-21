-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS events(
    id SERIAL NOT NULL,
    name TEXT NOT NULL,
    date DATE NOT NULL,
    info TEXT NOT NULL,
    category TEXT,

    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS events;
-- +goose StatementEnd
