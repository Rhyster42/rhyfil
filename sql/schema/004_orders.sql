-- +goose Up
CREATE TABLE orders (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    total DECIMAL
);

CREATE TABLE order_items (
    id UUID PRIMARY KEY,
    order_id UUID,
    CONSTRAINT fk_order_id
        FOREIGN KEY (order_id)
        REFERENCES orders(id),
    product_id UUID,
    CONSTRAINT fk_product_id
        FOREIGN KEY (product_id)
        REFERENCES products(id),
    quantity INTEGER,
    price_at_purchase DECIMAL
);

CREATE TABLE order_items_modifiers (
    order_item_id UUID,
    modifier_option_id UUID,
    PRIMARY KEY (order_item_id, modifier_option_id),
    CONSTRAINT fk_order_item_id
        FOREIGN KEY (order_item_id)
        REFERENCES order_items(id),
    CONSTRAINT fk_modifier_option_id
        FOREIGN KEY (modifier_option_id)
        REFERENCES modifier_options(id),
    price_adjustment_at_purchase DECIMAL
);

-- +goose Down
DROP TABLE order_items_modifiers;
DROP TABLE order_items;
DROP TABLE orders;