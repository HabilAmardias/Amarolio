CREATE TABLE IF NOT EXISTS transaction_schedulers (
    id TEXT PRIMARY KEY NOT NULL,

    frequency TEXT NOT NULL
        CHECK (frequency IN ('monthly', 'daily')),

    date_of_month INTEGER, -- monthly only, 1-28

    timezone TEXT NOT NULL,
    local_time TEXT NOT NULL,

    start_at TEXT NOT NULL,
    end_at TEXT,

    amount INTEGER NOT NULL,
    user_id TEXT NOT NULL,
    description TEXT NOT NULL,

    next_run_at TEXT NOT NULL,

    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    deleted_at TEXT,

    CONSTRAINT chk_date_of_month CHECK (
        (frequency = 'monthly' AND date_of_month BETWEEN 1 AND 28)
        OR
        (frequency = 'daily' AND date_of_month IS NULL)
    ),

    CONSTRAINT chk_date_range CHECK (
        end_at IS NULL OR end_at >= start_at
    )
);

CREATE INDEX IF NOT EXISTS idx_transaction_schedulers_user_id
    ON transaction_schedulers(user_id);

CREATE INDEX IF NOT EXISTS idx_transaction_schedulers_next_run_at
    ON transaction_schedulers(next_run_at)
    WHERE deleted_at IS NULL;
