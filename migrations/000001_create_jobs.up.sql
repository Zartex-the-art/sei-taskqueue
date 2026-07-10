
-- 001_jobs.sql
-- Canonical schema for Project 2

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE jobs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    queue_name TEXT NOT NULL DEFAULT 'default',

    payload JSONB NOT NULL,

    job_type TEXT NOT NULL,

    status TEXT NOT NULL DEFAULT 'queued'
        CHECK (status IN (
            'queued',
            'running',
            'done',
            'failed',
            'dead'
        )),

    priority INTEGER NOT NULL DEFAULT 0,

    attempts INTEGER NOT NULL DEFAULT 0,

    max_attempts INTEGER NOT NULL DEFAULT 5,

    run_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    locked_by TEXT,

    locked_until TIMESTAMPTZ,

    idempotency_key TEXT,

    last_error TEXT,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX idx_jobs_idempotency_key
ON jobs(idempotency_key)
WHERE idempotency_key IS NOT NULL;

CREATE INDEX idx_jobs_dequeue
ON jobs (status, run_at, priority DESC)
WHERE status = 'queued';

CREATE INDEX idx_jobs_locked_until
ON jobs(locked_until)
WHERE status = 'running';

CREATE TABLE dead_jobs (
    id UUID PRIMARY KEY,

    queue_name TEXT NOT NULL,

    payload JSONB NOT NULL,

    job_type TEXT NOT NULL,

    attempts INTEGER NOT NULL,

    last_error TEXT,

    original_created_at TIMESTAMPTZ NOT NULL,

    died_at TIMESTAMPTZ NOT NULL DEFAULT now()
);