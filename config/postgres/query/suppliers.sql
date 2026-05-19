-- name: CreateSupplier :one
INSERT INTO suppliers (name) VALUES ($1) RETURNING id;

-- name: GetSuppliers :many
SELECT * FROM suppliers ORDER BY name ASC;

-- name: GetSupplierById :one
SELECT * FROM suppliers WHERE id = $1 LIMIT 1;
