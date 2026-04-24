-- name: CreateProduct :one
INSERT INTO products (id, created_at, updated_at, name, count, price)
VALUES (
	$1,
	$2,
	$3,
	$4,
	$5,
	$6
)
RETURNING *;

-- name: DeleteAllProducts :exec
DELETE FROM products;


