-- +goose Up
CREATE TABLE orders (
    id UUID PRIMARY KEY NOT NULL,
    created_at TIMESTAMP NOT NULL,
    total DECIMAL NOT NULL
);

CREATE TABLE order_items (
    id UUID PRIMARY KEY NOT NULL,
    order_id UUID NOT NULL,
    CONSTRAINT fk_order_id
        FOREIGN KEY (order_id)
        REFERENCES orders(id),
    product_id UUID NOT NULL,
    CONSTRAINT fk_product_id
        FOREIGN KEY (product_id)
        REFERENCES products(id),
    quantity INTEGER NOT NULL,
    price_at_purchase DECIMAL NOT NULL
);

CREATE TABLE order_items_modifiers (
    order_item_id UUID NOT NULL,
    modifier_option_id UUID NOT NULL,
    PRIMARY KEY (order_item_id, modifier_option_id),
    CONSTRAINT fk_order_item_id
        FOREIGN KEY (order_item_id)
        REFERENCES order_items(id),
    CONSTRAINT fk_modifier_option_id
        FOREIGN KEY (modifier_option_id)
        REFERENCES modifier_options(id),
    price_adjustment_at_purchase DECIMAL NOT NULL
);

-- +goose Down
DROP TABLE order_items_modifiers;
DROP TABLE order_items;
DROP TABLE orders;