-- name: CreateUser :one
INSERT INTO users (
    email,
    username,

    salt,
    auth_verifier,

    identity_public_key,
    enc_identity_private_key,
    identity_private_key_nonce,

    vault_public_key,
    enc_vault_private_key,
    vault_private_key_nonce
) VALUES (
    $1, $2, $3, $4, 
    $5, $6, $7, 
    $8, $9, $10
)
RETURNING *;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetSaltByEmail :one
SELECT salt FROM users
WHERE email = $1 LIMIT 1;
