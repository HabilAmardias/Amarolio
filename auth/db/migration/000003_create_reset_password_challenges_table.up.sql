CREATE TABLE IF NOT EXISTS reset_password_challenges (
    id VARCHAR PRIMARY KEY NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    reset_password_challenge_expired_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);
