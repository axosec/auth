CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) NOT NULL UNIQUE,
    email_hash BYTEA NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,

    salt BYTEA NOT NULL,
    auth_verifier BYTEA NOT NULL,

    identity_public_key BYTEA NOT NULL,
    enc_identity_private_key BYTEA NOT NULL,
    identity_private_key_nonce BYTEA NOT NULL,

    vault_public_key BYTEA NOT NULL,
    enc_vault_private_key BYTEA NOT NULL,
    vault_private_key_nonce BYTEA NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
