CREATE TABLE IF NOT EXISTS product_prices (
  id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  product_id uuid NOT NULL REFERENCES products(id),
  price integer NOT NULL,
  effective_date timestamp NOT NULL DEFAULT (NOW()),
  created_at timestamp NOT NULL DEFAULT (NOW())
);
