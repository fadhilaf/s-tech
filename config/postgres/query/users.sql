-- name: CreateUser :execresult
INSERT INTO users (
  name, email, password_hash, address, phone
) VALUES ( 
  $1, $2, $3, $4, $5
);

-- name: GetUser :many
SELECT id, name, email FROM users
ORDER BY name;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
where email = $1 LIMIT 1;

-- name: UpdateUser :execresult
UPDATE users SET
  name = $2,
  address = $3,
  phone = $4
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
