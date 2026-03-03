-- name: CreateOrder :execresult
INSERT INTO orders (
  user_id, product_id, quantity, description
) VALUES (
  $1, $2, $3, $4
);

-- name: GetOrders :many
SELECT * FROM orders 
ORDER BY created_at DESC;

-- name: GetOrderById :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1;

-- name: GetOrdersByUserId :many
SELECT * FROM orders
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: UpdateOrder :execresult
UPDATE orders SET
  status = $2,
  description = $3
WHERE id = $1;

-- name: UpdateOrderStatus :execresult
UPDATE orders SET
  status = $2
WHERE id = $1;

-- name: DeleteOrder :exec 
DELETE FROM orders
WHERE id = $1;
