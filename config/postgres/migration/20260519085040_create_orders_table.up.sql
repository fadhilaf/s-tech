DO $$
  BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'order_status') THEN
      CREATE TYPE order_status AS ENUM ('pending', 'processing', 'delivered');
    END IF;
  END
$$;

CREATE TABLE IF NOT EXISTS orders (
  id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  user_id uuid NOT NULL REFERENCES users(id),
  product_price_id uuid NOT NULL REFERENCES product_prices(id),
  quantity integer NOT NULL,
  status order_status DEFAULT 'pending' NOT NULL,
  description varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT (NOW())
);
