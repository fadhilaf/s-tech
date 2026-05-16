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
