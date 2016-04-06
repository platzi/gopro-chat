
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO users (username, email, password, updated_at)
VALUES ('ivan','ivan@iver.mx','platzi', now());

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
TRUNCATE TABLE users RESTART IDENTITY;
