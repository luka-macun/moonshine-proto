.PHONY: generate run docker-up docker-down migrate-up migrate-down tidy

# Generate protobuf and ConnectRPC code
generate:
	buf lint
	buf generate

# Lint protos and check for backward-incompatible changes against main branch
lint:
	buf lint
	buf breaking --against '.git#branch=main'

# Run the server
run:
	go run ./cmd/server/main.go

# Start PostgreSQL via Docker Compose
docker-up:
	docker compose up -d

# Stop PostgreSQL
docker-down:
	docker compose down

# Run database migrations (up) via Docker
migrate-up:
	@echo "Applying migrations..."
	docker exec -i moonshine-postgres psql -U moonshine -d moonshine < internal/db/migrations/001_init.up.sql

# Roll back database migrations via Docker
migrate-down:
	@echo "Rolling back migrations..."
	docker exec -i moonshine-postgres psql -U moonshine -d moonshine < internal/db/migrations/001_init.down.sql

# Tidy Go modules
tidy:
	go mod tidy
