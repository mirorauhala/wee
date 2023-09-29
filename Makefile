.PHONY: run start-db stop-db migrate build clean

all: run

run:
	@echo "Running server..."
	@go run main.go

dev:
	@echo "Running server..."
	@air main.go

start-db:
	@echo "Starting database..."
	@docker compose up -d db

stop-db:
	@echo "Stopping server..."
	@docker compose down

migrate:
	@echo "Migrating database..."
	docker run -v ./migrations:/migrations --network host migrate/migrate:4 -path=/migrations/ -database postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable $(MIGRATE_COMMAND)

build:
	@echo "Building server..."
	@go build -o bin/server main.go

clean:
	@echo "Cleaning server..."
	@rm -rf bin
	@docker compose down --rmi all --volumes --remove-orphans
