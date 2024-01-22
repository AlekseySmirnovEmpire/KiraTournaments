-- +goose Up
-- +goose StatementBegin
CREATE TABLE users 
(
    id SERIAL NOT NULL UNIQUE,
    name varchar(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
