-- +goose Up
ALTER TABLE users ADD COLUMN
    hashed_password TEXT NOT NULL DEFAULT '';
ALTER TABLE users ALTER COLUMN hashed_password DROP DEFAULT;

-- +goose Down
ALTER TABLE users DROP COLUMN hashed_password;

