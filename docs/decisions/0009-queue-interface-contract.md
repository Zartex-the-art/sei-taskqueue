# ADR-009: Queue Interface Contract

## Status
Accepted

## Context
Every phase of this project — worker pool, HTTP API, admin
endpoints, future batching — needs a single stable contract to call
against, so no caller talks to Postgres directly. Getting this
wrong means touching every call site later instead of one
interface file.

## Decision
`Queue` exposes four methods: `Enqueue`, `Dequeue`, `Ack`, `Nack`.
Key choices and why:
**Dequeue returns a single job, not a batch.**
Keeps the worker pool (Day 8) and visibility timeout (Day 10) simple
to reason about. Batch claiming is a known lever for throughput
(flagged explicitly in the Day 21/22 tuning phase) but will be
added as an additional method later, if benchmarks show it's
needed — not baked into this signature speculatively now.

**Empty queue returns (nil, nil), not a sentinel error.**
A worker finding no work is the normal steady state, not a failure.
Treating it as an error would force every call site to unwrap an
`ErrNoJobs`-style check instead of a simple nil check.
**Ack/Nack take a job ID, not the full Job struct.**
Acknowledging completion doesn't require the caller to hold onto
the payload — and it keeps the interface's data requirements
honest.

## Consequences
Good:- One place to change if the claiming strategy evolves (e.g. batch
  claims added later)
  - Worker pool code stays simple 
  — no batch-vs-single branching logic
  needed until it's actually needed
  - Easy to write a fully in-memory implementation (Day 3) for fast
  unit tests, since the contract is small

Bad:- Single-job claims mean more round trips under very high load —
  explicitly deferred to Phase 6 benchmarking rather than solved
  speculatively today
