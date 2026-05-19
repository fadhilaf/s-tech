-- name: InsertProductStock :execresult
INSERT INTO product_stocks (
  product_id, supplier_id, is_add, quantity, price
) VALUES (
  $1, $2, $3, $4, $5
);

-- name: GetProductStocksByProductId :many
SELECT * FROM product_stocks WHERE product_id = $1 ORDER BY created_at DESC;
