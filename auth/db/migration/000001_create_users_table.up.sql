CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    email VARCHAR NOT NULL UNIQUE,
    password VARCHAR NOT NULL,
    otp VARCHAR,
    verified BOOLEAN NOT NULL DEFAULT FALSE,
    verification_token VARCHAR,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    otp_expired_at TIMESTAMPTZ,
    verification_token_expired_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);
