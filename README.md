# Moonshine Server

Backend server for the Moonshine chat app — a Discord/Slack-style communication platform. Built with **Go**, **ConnectRPC**, and **PostgreSQL**.

## Prerequisites

| Tool | Version | Install |
|------|---------|---------|
| **Go** | 1.24+ | [go.dev/dl](https://go.dev/dl/) |
| **Buf** | latest | `go install github.com/bufbuild/buf/cmd/buf@latest` |
| **protoc-gen-go** | latest | `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest` |
| **protoc-gen-connect-go** | latest | `go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest` |
| **Docker** | 20+ | [docker.com](https://www.docker.com/get-started/) |

> Make sure `$(go env GOPATH)/bin` is on your `PATH` so that `buf`, `protoc-gen-go`, and `protoc-gen-connect-go` are available.

## Go Dependencies

| Package | Purpose |
|---------|---------|
| [`connectrpc.com/connect`](https://connectrpc.com) | ConnectRPC framework (HTTP API layer) |
| [`github.com/jackc/pgx/v5`](https://github.com/jackc/pgx) | PostgreSQL driver & connection pool |
| [`google.golang.org/protobuf`](https://pkg.go.dev/google.golang.org/protobuf) | Protocol Buffers runtime |
| [`buf.build/gen/go/bufbuild/protovalidate`](https://buf.build/bufbuild/protovalidate) | Request validation via proto annotations |

## Quick Start

```bash
# 1. Start PostgreSQL
make docker-up

# 2. Apply database migrations
make migrate-up

# 3. Install Go dependencies
go mod tidy

# 4. Run the server
make run
# → Moonshine server listening on :8080
```

## Available Make Targets

| Command | Description |
|---------|-------------|
| `make generate` | Regenerate protobuf & ConnectRPC code |
| `make run` | Start the server |
| `make docker-up` | Start PostgreSQL container |
| `make docker-down` | Stop PostgreSQL container |
| `make migrate-up` | Apply database migrations |
| `make migrate-down` | Roll back database migrations |
| `make tidy` | Run `go mod tidy` |

## Project Structure

```
├── proto/moonshine/v1/       # Protobuf service definitions
├── gen/                      # Generated Go code (via buf generate)
├── cmd/server/               # Server entrypoint
├── internal/
│   ├── config/               # Environment-based configuration
│   ├── db/                   # Database connection & migrations
│   └── service/              # ConnectRPC service handler stubs
├── docker-compose.yml        # PostgreSQL 16
├── buf.yaml / buf.gen.yaml   # Buf codegen config
└── Makefile
```

## Configuration

Set via environment variables or copy `.env.example` to `.env`:

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | Server listen port |
| `DATABASE_URL` | `postgres://moonshine:moonshine@localhost:5432/moonshine?sslmode=disable` | PostgreSQL connection string |
| `JWT_SECRET` | `dev-secret-change-me` | Secret for signing JWT tokens |
