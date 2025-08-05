CREATE TABLE IF NOT EXISTS backlite_tasks (
    id text PRIMARY KEY,
    created_at BIGINT NOT NULL,
    queue text NOT NULL,
    task BYTEA NOT NULL,
    wait_until BIGINT,
    claimed_at BIGINT,
    last_executed_at BIGINT,
    attempts BIGINT NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS backlite_tasks_completed (
    id text PRIMARY KEY NOT NULL,
    created_at BIGINT NOT NULL,
    queue text NOT NULL,
    last_executed_at BIGINT,
    attempts BIGINT NOT NULL,
    last_duration_micro BIGINT,
    succeeded BIGINT,
    task BYTEA,
    expires_at BIGINT,
    error text
);

CREATE INDEX IF NOT EXISTS backlite_tasks_wait_until ON backlite_tasks (wait_until) WHERE wait_until IS NOT NULL;
