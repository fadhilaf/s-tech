CREATE TABLE IF NOT EXISTS product_stocks (
  id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  product_id uuid NOT NULL REFERENCES products(id),
  supplier_id uuid NOT NULL REFERENCES suppliers(id),
  is_add boolean NOT NULL,
  quantity integer NOT NULL,
  price integer NOT NULL,
  created_at timestamp NOT NULL DEFAULT (NOW())
);

-- Trigger to update product stock
CREATE OR REPLACE FUNCTION update_product_stock_trigger()
RETURNS TRIGGER AS $$
BEGIN
  IF TG_OP = 'INSERT' THEN
    IF NEW.is_add THEN
      UPDATE products SET stock = stock + NEW.quantity WHERE id = NEW.product_id;
    ELSE
      UPDATE products SET stock = stock - NEW.quantity WHERE id = NEW.product_id;
    END IF;
  ELSIF TG_OP = 'DELETE' THEN
    IF OLD.is_add THEN
      UPDATE products SET stock = stock - OLD.quantity WHERE id = OLD.product_id;
    ELSE
      UPDATE products SET stock = stock + OLD.quantity WHERE id = OLD.product_id;
    END IF;
  ELSIF TG_OP = 'UPDATE' THEN
    -- Revert old
    IF OLD.is_add THEN
      UPDATE products SET stock = stock - OLD.quantity WHERE id = OLD.product_id;
    ELSE
      UPDATE products SET stock = stock + OLD.quantity WHERE id = OLD.product_id;
    END IF;
    -- Apply new
    IF NEW.is_add THEN
      UPDATE products SET stock = stock + NEW.quantity WHERE id = NEW.product_id;
    ELSE
      UPDATE products SET stock = stock - NEW.quantity WHERE id = NEW.product_id;
    END IF;
  END IF;
  RETURN NULL;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_product_stock
AFTER INSERT OR UPDATE OR DELETE ON product_stocks
FOR EACH ROW EXECUTE FUNCTION update_product_stock_trigger();
