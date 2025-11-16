.PHONY: help dev build run docker-up docker-down migrate-create migrate-up migrate-down migrate-status migrate-validate db-reset test clean

# Load environment variables from .env file
include .env
export

# Default target
help:
	@echo "Available commands:"
	@echo "  make dev              - Run development server with hot reload"
	@echo "  make build            - Build the application"
	@echo "  make run              - Run the application"
	@echo "  make docker-up        - Start docker containers"
	@echo "  make docker-down      - Stop docker containers"
	@echo "  make migrate-create   - Create a new migration (name=migration_name)"
	@echo "  make migrate-up       - Apply all pending migrations"
	@echo "  make migrate-down     - Rollback last migration"
	@echo "  make migrate-status   - Show migration status"
	@echo "  make migrate-validate - Validate migrations"
	@echo "  make db-reset         - Reset database and reapply migrations"
	@echo "  make test             - Run tests"
	@echo "  make clean            - Clean build artifacts"

# Development
dev:
	air

# Build
build:
	go build -o tmp/main cmd/server/main.go

# Run
run: build
	./tmp/main

# Docker
docker-up:
	docker-compose -f docker-compose.dev.yml up -d

docker-down:
	docker-compose -f docker-compose.dev.yml down

# Atlas Migrations
migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "Error: Please provide a migration name. Usage: make migrate-create name=your_migration_name"; \
		exit 1; \
	fi
	atlas migrate diff $(name) --env local

migrate-up:
	atlas migrate apply --env local

migrate-down:
	atlas migrate down --env local

migrate-status:
	atlas migrate status --env local

migrate-validate:
	atlas migrate validate --env local

migrate-hash:
	atlas migrate hash --env local

# Database operations
db-reset:
	@echo "Dropping all tables in users schema..."
	docker exec -it ayushchugh.com-postgres psql -U postgres -d ayushchugh.com -c "DROP SCHEMA IF EXISTS users CASCADE;"
	docker exec -it ayushchugh.com-postgres psql -U postgres -d ayushchugh.com -c "DROP TABLE IF EXISTS atlas_schema_revisions CASCADE;"
	@echo "Reapplying migrations..."
	$(MAKE) migrate-up

# Testing
test:
	go test -v ./...

# Clean
clean:
	rm -rf tmp/
	go clean