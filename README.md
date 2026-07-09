# Task Queue Engine — Project 2
> Status: Phase 1 (Foundation) — in progress

## What This Is
A durable, observable, distributed task queue built on Postgres +
Redis, sustaining 5,000–10,000 jobs/min with at-least-once delivery
and crash recovery.

## Quick Start
(Filled in once `docker compose up` boots the full stack — Day 4 gate)

## Architecture
(Diagram + explanation — Day 4 onward)

The system uses a Queue Interface to separate workers from the underlying PostgreSQL implementation. Workers interact only with the queue interface to dequeue jobs and acknowledge completed work, while the queue interface manages database operations internally. The design decision is documented in ADR-009.

See the sequence diagram for the interaction flow:

- Queue Interface Flow: `docs/diagrams/queue-interface-flow.png`
- ADR-009: `docs/decisions/0009-queue-interface.md`

## API Reference
(Filled in starting Day 7)

## Reliability Guarantees
(Filled in starting Day 10 — visibility timeout, retries, DLQ)

## Scheduling & Priority
(Filled in starting Day 15)

## Observability
(Filled in starting Day 18 — metrics, dashboard)

## Failure Modes
(Filled in Day 14 and Phase 6)

## Benchmarks
(Filled in Days 21–22)

## Architecture Decision Records
See `docs/decisions/` — index maintained here as ADRs land.

## What We'd Do at 10x Scale
(Filled in Day 23)


# System Container Map – v1 (Day 3 draft)

This is a forward-looking diagram showing the architecture expected after Day 4, where the application, PostgreSQL, and Redis run together using Docker Compose.

Currently only PostgreSQL and Redis are available in Compose. The application container image was built on Day 3 and will join the network on Day 4.

PostgreSQL is the system of record. Redis is present, but its final role is intentionally left undecided and will be finalized later in the project.