-- +goose Up
CREATE TABLE products (
	id UUID PRIMARY KEY,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	name TEXT NOT NULL,
	count INTEGER,
	price DECIMAL
);

-- +goose Down
DROP TABLE products;
