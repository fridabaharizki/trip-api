# trip-api

Simple REST API for managing trips, built with Go (stdlib only).

## Features

- CRUD endpoints for trips
- In-memory store behind a `Store` interface (Postgres planned)
- Structured logging (`log/slog`), graceful shutdown
- Table-driven tests, multi-stage Docker build

## Run

```bash
go run ./cmd/api
# or
docker build -t trip-api . && docker run -p 8080:8080 trip-api
```

## API

| Method | Path          | Description  |
| ------ | ------------- | ------------ |
| GET    | `/healthz`    | Health check |
| GET    | `/trips`      | List trips   |
| POST   | `/trips`      | Create trip  |
| GET    | `/trips/{id}` | Get trip     |
| DELETE | `/trips/{id}` | Delete trip  |

```bash
curl -X POST localhost:8080/trips -d '{"destination":"Bali"}'
```

## Test

```bash
go test ./... -race
```

## Roadmap

- [ ] PostgreSQL (pgx + sqlc)
- [ ] docker-compose
- [ ] GitHub Actions CI
- [ ] React + TypeScript frontend
