
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users
(
  id SERIAL PRIMARY KEY,
  username VARCHAR(64),
  email VARCHAR(80),
  password VARCHAR(80),
  created_at TIMESTAMP DEFAULT now() NOT NULL,
  updated_at TIMESTAMP NOT NULL
);
CREATE UNIQUE INDEX unique_username ON users (username);
CREATE UNIQUE INDEX unique_email ON users (email);
CREATE INDEX ix_users_id ON users (id);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;
