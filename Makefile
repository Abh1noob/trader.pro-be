# Makefile for trader.pro backend

# Variables
APP_NAME = trader.pro-be
PORT = 8080

# Default target
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make run         - Run the application with go run"
	@echo "  make build       - Build the application binary"
	@echo "  make migrate     - Run database migrations"
	@echo "  make air         - Run the application with Air for live reloading"
	@echo "  make test        - Run tests"
	@echo "  make clean       - Remove built binaries"

# Run the application directly with go run
.PHONY: run
run:
	go run main.go

# Build the application binary
.PHONY: build
build:
	go build -o $(APP_NAME) main.go

# Run database migrations
.PHONY: migrate
migrate:
	go run cmd/migrations/migrate.go

# Run the application with Air for live reloading
.PHONY: air
air:
	air

# Run tests
.PHONY: test
test:
	go test ./... -v

# Clean up built binaries
.PHONY: clean
clean:
	rm -f $(APP_NAME)