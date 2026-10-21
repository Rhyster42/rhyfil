# rhyfil

Point-of-sale web app for tablets, built for coffee-shop workflows: product → modifiers → transaction.

## Stack
- Go (net/http), PostgreSQL via lib/pq
- sqlc for query generation, goose for migrations
- Vanilla JS/HTML/CSS frontend

## Layout
- `cmd/` — `main` package: CLI entrypoint (`main.go`), command registry (`commands.go`), and `handler_*.go` for CLI commands and HTTP handlers (`handler_serve.go` registers routes)
- `server/` — JSON response helpers
- `internal/config/` — reads `~/.rhyfilconfig.json` (`db_url`, `current_user_name`)
- `internal/database/` — sqlc-generated code; do not edit by hand
- `sql/schema/` — goose migrations; `sql/queries/` — sqlc queries
- `frontend/` — static files served at `/`

## Workflow
- After changing `sql/queries/` or `sql/schema/`, run `sqlc generate` from the repo root
- Migrations: `goose -dir sql/schema postgres "<db_url>" up`
- `serve` serves `../frontend` relative to the working directory, so run it from `cmd/`
- CLI commands are registered in `getCommands()` in `cmd/commands.go`
