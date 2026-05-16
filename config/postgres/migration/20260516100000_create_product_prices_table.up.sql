CREATE TABLE IF NOT EXISTS product_prices (
    id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    product_id uuid NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    price integer NOT NULL,
    effective_date timestamp NOT NULL DEFAULT (NOW()),
    created_at timestamp NOT NULL DEFAULT (NOW())
);

CREATE INDEX IF NOT EXISTS product_prices_product_id_idx ON product_prices (product_id);
