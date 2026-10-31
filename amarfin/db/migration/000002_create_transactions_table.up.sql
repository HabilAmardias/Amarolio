
CREATE TABLE IF NOT EXISTS transactions (
    id TEXT PRIMARY KEY NOT NULL,
    amount INTEGER NOT NULL,
    user_id TEXT NOT NULL,
    description TEXT NOT NULL,

    scheduler_id TEXT,
    scheduled_for TEXT, -- timestamptz

    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    deleted_at TEXT,

    FOREIGN KEY (scheduler_id)
        REFERENCES transaction_schedulers(id)
        ON DELETE RESTRICT,

    CONSTRAINT uq_scheduler_execution
        UNIQUE (scheduler_id, scheduled_for),

    CONSTRAINT chk_scheduler_execution CHECK (
        (scheduler_id IS NULL AND scheduled_for IS NULL)
        OR
        (scheduler_id IS NOT NULL AND scheduled_for IS NOT NULL)
    )
);

CREATE INDEX IF NOT EXISTS idx_transactions_user_id
    ON transactions(user_id);

CREATE INDEX IF NOT EXISTS idx_transactions_scheduler_id
    ON transactions(scheduler_id);
