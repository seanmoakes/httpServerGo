-- +goose Up
ALTER TABLE users
  ADD COLUMN hashed_password TEXT NOT NULL;

UPDATE users
  SET hashed_password = 'unset';
-- +goose Down
