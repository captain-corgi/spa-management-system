# Database Makefile for local development

MIGRATIONS_DIR ?= ./migrations
DATABASE_DSN ?= postgres://postgres:postgres@localhost:5432/spa?sslmode=disable

migrate:
	@echo ">>> Running database migrations..."
	DATABASE_DSN=$(DATABASE_DSN) MIGRATIONS_DIR=$(MIGRATIONS_DIR) go run migrate.go

.PHONY: migrate
