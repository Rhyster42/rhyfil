-- name: CreateModifierGroup :one
INSERT INTO modifier_groups(id, name)
VALUES (
    $1,
    $2
)
RETURNING *;

-- name: CreateModifierOption :one
INSERT INTO modifier_options(id, name, price_adjustment, modifier_group_id)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: LinkModifierGroupToProduct :exec
INSERT INTO product_modifier_groups(product_id, modifier_group_id)
VALUES (
    $1,
    $2
);

-- name: GetAllModifierGroupsForProduct :many
SELECT modifier_groups.* FROM product_modifier_groups
INNER JOIN modifier_groups
    ON product_modifier_groups.modifier_group_id = modifier_groups.id
WHERE $1 = product_modifier_groups.product_id;

-- name: GetAllOptionsForModifierGroup :many
SELECT * FROM modifier_options
WHERE $1 = modifier_group_id;
