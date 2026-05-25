-- name: CreateProduct :one
INSERT INTO products (
  name, stock, is_service, description, image
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING id;

-- name: CreateProductPrice :execresult
INSERT INTO product_prices (
    product_id, price, effective_date
) VALUES (
    $1, $2, $3
);

-- name: GetProduct :many
SELECT
    p.id,
    p.name,
    p.stock,
    p.is_service,
    p.description,
    p.image,
    pp.id AS product_price_id,
    pp.price AS current_price
FROM products p
JOIN product_prices pp ON p.id = pp.product_id
WHERE pp.effective_date = (
    SELECT MAX(effective_date)
    FROM product_prices
    WHERE product_id = p.id
      AND effective_date <= NOW()
)
ORDER BY p.created_at DESC;

-- name: GetProductById :one
SELECT
    p.id,
    p.name,
    p.stock,
    p.is_service,
    p.description,
    p.image,
    pp.id AS product_price_id,
    pp.price AS current_price
FROM products p
JOIN product_prices pp ON p.id = pp.product_id
WHERE p.id = $1
  AND pp.effective_date = (
    SELECT MAX(effective_date)
    FROM product_prices
    WHERE product_id = p.id
      AND effective_date <= NOW()
  )
LIMIT 1;

-- name: GetProductByName :one
SELECT
    p.id,
    p.name,
    p.stock,
    p.is_service,
    p.description,
    p.image,
    pp.id AS product_price_id,
    pp.price AS current_price
FROM products p
JOIN product_prices pp ON p.id = pp.product_id
WHERE p.name = $1
  AND pp.effective_date = (
    SELECT MAX(effective_date)
    FROM product_prices
    WHERE product_id = p.id
      AND effective_date <= NOW()
  )
LIMIT 1;

-- name: GetProductByQuery :many
SELECT
    p.id,
    p.name,
    p.stock,
    p.is_service,
    p.description,
    p.image,
    pp.id AS product_price_id,
    pp.price AS current_price
FROM products p
JOIN product_prices pp ON p.id = pp.product_id
WHERE p.name ILIKE $1
  AND pp.effective_date = (
    SELECT MAX(effective_date)
    FROM product_prices
    WHERE product_id = p.id
      AND effective_date <= NOW()
  )
ORDER BY p.name;

-- name: UpdateProduct :execresult
UPDATE products SET
  name = $2,
  stock = $3,
  description = $4,
  image = $5
WHERE id = $1;

-- name: UpdateProductPrice :execresult
INSERT INTO product_prices (
    product_id, price, effective_date
) VALUES (
    $1, $2, NOW()
);

-- name: UpdateProductStock :execresult
UPDATE products SET
  stock = $2
WHERE id = $1;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;

-- name: GetProductByPriceId :one
SELECT
    p.id,
    p.name,
    p.stock,
    p.is_service,
    p.description,
    p.image,
    pp.id AS product_price_id,
    pp.price AS current_price
FROM products p
JOIN product_prices pp ON p.id = pp.product_id
WHERE pp.id = $1
LIMIT 1;

-- name: UpdateProductDetails :execresult
UPDATE products SET
  name = $2,
  description = $3,
  image = $4
WHERE id = $1;

-- name: GetProductChronology :many
SELECT
  'STOCK_IN'::varchar AS log_type,
  ps.created_at AS date,
  ps.quantity AS quantity,
  ps.price AS cost_price,
  0::integer AS sell_price,
  COALESCE(s.name, '') AS related_name
FROM product_stocks ps
LEFT JOIN suppliers s ON ps.supplier_id = s.id
WHERE ps.product_id = $1 AND ps.is_add = true

UNION ALL

SELECT
  'STOCK_OUT'::varchar AS log_type,
  o.created_at AS date,
  o.quantity AS quantity,
  0::integer AS cost_price,
  pp.price AS sell_price,
  COALESCE(u.name, '') AS related_name
FROM orders o
JOIN product_prices pp ON o.product_price_id = pp.id
JOIN users u ON o.user_id = u.id
WHERE pp.product_id = $1

ORDER BY date DESC;

-- name: GetProductPricesByProductId :many
SELECT
  id,
  product_id,
  price,
  effective_date,
  created_at
FROM product_prices
WHERE product_id = $1
ORDER BY effective_date ASC;

-- name: GetAllChronology :many
SELECT
  'STOCK_IN'::varchar AS log_type,
  ps.created_at AS date,
  ps.quantity AS quantity,
  ps.price AS cost_price,
  0::integer AS sell_price,
  COALESCE(s.name, '') AS related_name,
  p.name AS product_name
FROM product_stocks ps
LEFT JOIN suppliers s ON ps.supplier_id = s.id
JOIN products p ON ps.product_id = p.id
WHERE ps.is_add = true

UNION ALL

SELECT
  'STOCK_OUT'::varchar AS log_type,
  o.created_at AS date,
  o.quantity AS quantity,
  0::integer AS cost_price,
  pp.price AS sell_price,
  COALESCE(u.name, '') AS related_name,
  p.name AS product_name
FROM orders o
JOIN product_prices pp ON o.product_price_id = pp.id
JOIN users u ON o.user_id = u.id
JOIN products p ON pp.product_id = p.id

ORDER BY date DESC;
