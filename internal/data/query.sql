-- name: CreateUser :one
INSERT INTO users (
    email,
    salt,
    auth_verifier,
    public_key,
    enc_private_key
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetSaltByEmail :one
SELECT salt FROM users
WHERE email = $1 LIMIT 1;

-- name: GetLoginDetails :one
SELECT id, auth_verifier, enc_private_key, public_key
FROM users
WHERE email = $1 LIMIT 1;
