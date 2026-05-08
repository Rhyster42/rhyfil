-- name: CreateOrder :one
INSERT INTO orders(id, created_at, total)
VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: AddItemToOrder :one
INSERT INTO order_items (id, order_id, product_id, quantity, price_at_purchase)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;

-- name: AddModifierToOrderItem :one
INSERT INTO order_items_modifiers (order_item_id, modifier_option_id, price_adjustment_at_purchase)
VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetAllItemsInOrder :many
SELECT order_items.quantity, order_items.price_at_purchase, products.name
FROM order_items
INNER JOIN products
    ON order_items.product_id = products.id
WHERE order_items.order_id = $1;