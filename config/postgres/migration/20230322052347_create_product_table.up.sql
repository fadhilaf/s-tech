CREATE TABLE IF NOT EXISTS products (
  id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  name varchar(255) NOT NULL,
  price integer NOT NULL,
  stock integer NOT NULL,
  is_service boolean NOT NULL DEFAULT false,
  description varchar(255) NOT NULL,
  image varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT (NOW())
);
