# Core Ticketing Engine

A backend ticketing REST API built with Go, exploring production-grade 
patterns for data consistency, concurrent safety, and clean architecture 
used in high-traffic ticket platforms (e.g. concert/event ticketing).

## Architecture

Clean Architecture with layered separation:
- **Delivery** — HTTP handlers (stdlib net/http + ServeMux)
- **Use Case** — business logic
- **Repository** — data access via raw SQL (sqlx)
- **Domain** — core entities and interfaces

## Tech Stack

Go 1.25 · PostgreSQL · sqlx · Docker · Docker Compose

## What's Implemented

- Ticket creation and retrieval with Clean Architecture
- ACID-compliant transactions via sqlx to prevent data inconsistency
- Raw SQL with DTO pattern for type-safe data transfer
- Containerized PostgreSQL via Docker Compose

## Engineering Decisions

The central problem: what happens when 10,000 users hit 
"Buy" on the same ticket simultaneously?

| Failure Mode | Without Fix | Solution Implemented |
|---|---|---|
| Two users buy the last ticket at the same millisecond | Oversell — both get confirmation | Pessimistic locking (`SELECT FOR UPDATE`) — DB serializes concurrent buyers |
| Payment succeeds but reservation fails | User charged, no ticket received | ACID transaction — debit + reservation are atomic (all or nothing) |
| 10k users refresh ticket availability per second | PostgreSQL hammered with identical reads | Redis cache — serve from memory, DB only hit on cache miss |
| Raw DB object sent to client | Internal pricing logic / IDs exposed | DTO pattern — explicit data contract between domain and presentation |

## In Progress

- **Pessimistic locking** — `SELECT FOR UPDATE` at DB level to serialize 
  concurrent ticket purchases and prevent overselling
- **Redis caching layer** — in-memory read cache to reduce PostgreSQL I/O 
  on high-frequency ticket fetch scenarios

## Run Locally

```bash
docker-compose up -d     # start PostgreSQL
go run ./cmd/api         # start server
```

## Why This Project

Most Go tutorials stop at "build a REST API." This project intentionally 
goes further — exploring what happens when multiple users hit the same 
ticket simultaneously, and how to handle it correctly at the database level 
before adding any caching layer on top.