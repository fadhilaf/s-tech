CREATE TABLE IF NOT EXISTS suppliers (
  id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  name varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT (NOW())
);

-- Insert undocumented supplier with a constant UUID for easy reference
INSERT INTO suppliers (id, name) VALUES ('00000000-0000-0000-0000-000000000000', 'Undocumented') ON CONFLICT (id) DO NOTHING;
