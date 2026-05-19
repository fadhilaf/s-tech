CREATE TABLE IF NOT EXISTS users (
  id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  name varchar(255) NOT NULL,
  email varchar(255) UNIQUE NOT NULL,
  password_hash varchar(255) NOT NULL,
  address varchar(255) NOT NULL,
  phone varchar(20) NOT NULL,
  created_at timestamp NOT NULL DEFAULT (NOW())
);
