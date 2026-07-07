# ADR-008: Postgres for Job State

## Status
Accepted

## Context
Project 1 was Redis-only — every piece of state (counters, rules)
lived in Redis, and that was the right call for a rate limiter:
small values, high-frequency reads/writes, no need for complex
querying or joins.
A task queue is a different shape of problem. Jobs are structured
records with many fields (status, priority, attempts, timestamps)
that need to be queried, filtered, and ordered in combination — 	"give
me the oldest queued job in this priority tier that isn't locked.	"
That's a relational query, not a key-value lookup.

## Alternatives Considered
**Redis-only, using sorted sets / lists for queuing.**
Would work for simple FIFO, but modeling retries, dead-letter
routing, and priority + delay together as Redis data structures gets
complex fast, and durability guarantees are weaker than Postgres's
WAL-backed writes without extra Redis persistence configuration.

**Postgres only, no Redis.**
Viable — Postgres alone can do everything this project needs. Redis
stays in the stack from Project 1 mainly for future use (e.g.
pub/sub for reduced polling latency, to be decided explicitly by Day
14) rather than as a hard requirement of the queue's correctness.

## Decision
Postgres is the system of record for all job state. `SELECT ... FOR
UPDATE SKIP LOCKED` (Day 5) gives safe concurrent dequeue without a
separate locking layer, and a single relational table lets one query
answer 	"next job to run	" across status, run_at, and priority at
once — something that would require multiple Redis data structures
kept in sync by hand.

## Consequences
Good:
- One source of truth, one place to look for job state
- Native support for the exact multi-column queries this project needs
- Durability comes from Postgres's WAL, not extra Redis config

Bad:
- New database for the team (Days 1–4 explicitly budget for this)
- Postgres write throughput must be validated at 10K jobs/min
  (addressed in Phase 6 benchmarks)