-- +goose Up
CREATE TABLE modifier_groups (
	id UUID PRIMARY KEY,
	name TEXT NOT NULL
);

CREATE TABLE modifier_options (
	id UUID PRIMARY KEY,
	name TEXT NOT NULL,
	price_adjustment DECIMAL NOT NULL,
	modifier_group_id UUID NOT NULL,
	CONSTRAINT fk_modifier_groups
		FOREIGN KEY (modifier_group_id)
		REFERENCES modifier_groups(id)
);

CREATE TABLE product_modifier_groups (
	product_id UUID,
	modifier_group_id UUID,
	PRIMARY KEY (product_id, modifier_group_id),
	CONSTRAINT fk_product
		FOREIGN KEY (product_id)
		REFERENCES products(id),
	CONSTRAINT fk_modifier_group
		FOREIGN KEY (modifier_group_id)
		REFERENCES modifier_groups(id)
);

-- +goose Down

DROP TABLE product_modifier_groups;
DROP TABLE modifier_options;
DROP TABLE modifier_groups;
