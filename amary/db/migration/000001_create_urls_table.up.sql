CREATE TABLE IF NOT EXISTS urls (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    user_id UUID, -- can be null if user is not logged-in
    long_url VARCHAR NOT NULL,
    short_code VARCHAR,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    expired_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_urls_short_code ON urls(short_code);
