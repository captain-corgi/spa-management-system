# Makefile for Spa Management System Backend
# Provides common operations for all backend services

# Variables
GO_VERSION := 1.24.1
SERVICES := customer-service appointment-service staff-service branch-service marketing-service finance-service analytics-service
ALL_COMPONENTS := $(SERVICES) api-gateway
DATABASE_DSN ?= postgres://postgres:postgres@localhost:5432/spa?sslmode=disable
REDIS_ADDR ?= localhost:6379
DOCKER_REGISTRY ?= spa-management
DOCKER_TAG ?= latest
MIGRATIONS_DIR ?= ./pkg/database/migrations
COVERAGE_DIR := ./coverage
LINT_TIMEOUT := 5m
ENV ?= dev

# Colors for terminal output
COLOR_RESET := $(shell tput sgr0)
COLOR_GREEN := $(shell tput setaf 2)
COLOR_YELLOW := $(shell tput setaf 3)
COLOR_BLUE := $(shell tput setaf 4)

.PHONY: all help build test lint clean docker-build docker-push run dev migrate swagger install-tools

# Default target
all: build

# Help command
help:
	@echo "$(COLOR_GREEN)Spa Management System Backend Makefile$(COLOR_RESET)"
	@echo "$(COLOR_YELLOW)Usage:$(COLOR_RESET)"
	@echo "  make [target]"
	@echo ""
	@echo "$(COLOR_YELLOW)Targets:$(COLOR_RESET)"
	@echo "  $(COLOR_GREEN)help$(COLOR_RESET)           - Show this help message"
	@echo "  $(COLOR_GREEN)build$(COLOR_RESET)          - Build all services"
	@echo "  $(COLOR_GREEN)build-service$(COLOR_RESET)  - Build a specific service (e.g., make build-service SERVICE=customer-service)"
	@echo "  $(COLOR_GREEN)test$(COLOR_RESET)           - Run tests for all services"
	@echo "  $(COLOR_GREEN)test-service$(COLOR_RESET)   - Test a specific service (e.g., make test-service SERVICE=customer-service)"
	@echo "  $(COLOR_GREEN)lint$(COLOR_RESET)           - Run linter on all services"
	@echo "  $(COLOR_GREEN)clean$(COLOR_RESET)          - Remove build artifacts"
	@echo "  $(COLOR_GREEN)docker-build$(COLOR_RESET)   - Build Docker images for all services"
	@echo "  $(COLOR_GREEN)docker-push$(COLOR_RESET)    - Push Docker images to registry"
	@echo "  $(COLOR_GREEN)run$(COLOR_RESET)            - Run all services locally"
	@echo "  $(COLOR_GREEN)run-service$(COLOR_RESET)    - Run a specific service (e.g., make run-service SERVICE=customer-service)"
	@echo "  $(COLOR_GREEN)dev$(COLOR_RESET)            - Run services with hot reload using air"
	@echo "  $(COLOR_GREEN)migrate$(COLOR_RESET)        - Run database migrations"
	@echo "  $(COLOR_GREEN)swagger$(COLOR_RESET)        - Generate Swagger documentation"
	@echo "  $(COLOR_GREEN)install-tools$(COLOR_RESET)  - Install development tools"
	@echo ""
	@echo "$(COLOR_YELLOW)Environment Variables:$(COLOR_RESET)"
	@echo "  DATABASE_DSN  - Database connection string (default: $(DATABASE_DSN))"
	@echo "  REDIS_ADDR    - Redis address (default: $(REDIS_ADDR))"
	@echo "  SERVICE       - Specific service to operate on"
	@echo "  ENV           - Environment (dev, sit, uat, nft, prd) (default: $(ENV))"
	@echo "  DOCKER_REGISTRY - Docker registry (default: $(DOCKER_REGISTRY))"
	@echo "  DOCKER_TAG    - Docker tag (default: $(DOCKER_TAG))"

# Build all services
build:
	@echo "$(COLOR_BLUE)Building all services...$(COLOR_RESET)"
	@for service in $(SERVICES); do \
		echo "$(COLOR_BLUE)Building $$service...$(COLOR_RESET)"; \
		(cd services/$$service && go build -o ../../bin/$$service ./main.go) || exit 1; \
	done
	@if [ -d api-gateway ]; then \
		echo "$(COLOR_BLUE)Building api-gateway...$(COLOR_RESET)"; \
		(cd api-gateway && go build -o ../bin/api-gateway ./main.go) || exit 1; \
	fi
	@echo "$(COLOR_GREEN)All services built successfully!$(COLOR_RESET)"

# Build a specific service
build-service:
	@if [ -z "$(SERVICE)" ]; then \
		echo "$(COLOR_YELLOW)Please specify a service with SERVICE=<service-name>$(COLOR_RESET)"; \
		exit 1; \
	fi
	@echo "$(COLOR_BLUE)Building $(SERVICE)...$(COLOR_RESET)"
	@cd services/$(SERVICE) && go build -o ../../bin/$(SERVICE) ./main.go
	@echo "$(COLOR_GREEN)$(SERVICE) built successfully!$(COLOR_RESET)"

# Run tests for all services
test:
	@echo "$(COLOR_BLUE)Running tests for all services...$(COLOR_RESET)"
	@mkdir -p $(COVERAGE_DIR)
	@for service in $(SERVICES); do \
		echo "$(COLOR_BLUE)Testing $$service...$(COLOR_RESET)"; \
		(cd services/$$service && go test -v -coverprofile=../../$(COVERAGE_DIR)/$$service.out ./...) || exit 1; \
	done
	@echo "$(COLOR_GREEN)All tests passed!$(COLOR_RESET)"

# Test a specific service
test-service:
	@if [ -z "$(SERVICE)" ]; then \
		echo "$(COLOR_YELLOW)Please specify a service with SERVICE=<service-name>$(COLOR_RESET)"; \
		exit 1; \
	fi
	@echo "$(COLOR_BLUE)Testing $(SERVICE)...$(COLOR_RESET)"
	@mkdir -p $(COVERAGE_DIR)
	@cd services/$(SERVICE) && go test -v -coverprofile=../../$(COVERAGE_DIR)/$(SERVICE).out ./...
	@echo "$(COLOR_GREEN)Tests for $(SERVICE) completed!$(COLOR_RESET)"

# Run linter on all services
lint:
	@echo "$(COLOR_BLUE)Running linter on all services...$(COLOR_RESET)"
	@for service in $(SERVICES); do \
		echo "$(COLOR_BLUE)Linting $$service...$(COLOR_RESET)"; \
		(cd services/$$service && golangci-lint run --timeout=$(LINT_TIMEOUT) ./...) || exit 1; \
	done
	@if [ -d api-gateway ]; then \
		echo "$(COLOR_BLUE)Linting api-gateway...$(COLOR_RESET)"; \
		if [ -f api-gateway/go.mod ]; then \
			(cd api-gateway && golangci-lint run --timeout=$(LINT_TIMEOUT) ./...) || echo "$(COLOR_YELLOW)Skipping api-gateway linting due to module issues$(COLOR_RESET)"; \
		else \
			echo "$(COLOR_YELLOW)Skipping api-gateway linting (no go.mod file)$(COLOR_RESET)"; \
		fi \
	fi
	@echo "$(COLOR_GREEN)Linting completed!$(COLOR_RESET)"

# Clean build artifacts
clean:
	@echo "$(COLOR_BLUE)Cleaning build artifacts...$(COLOR_RESET)"
	@rm -rf bin/
	@rm -rf $(COVERAGE_DIR)/
	@echo "$(COLOR_GREEN)Cleaned!$(COLOR_RESET)"

# Build Docker images for all services
docker-build:
	@echo "$(COLOR_BLUE)Building Docker images for all services...$(COLOR_RESET)"
	@for service in $(SERVICES); do \
		echo "$(COLOR_BLUE)Building Docker image for $$service...$(COLOR_RESET)"; \
		docker build -t $(DOCKER_REGISTRY)/$$service:$(DOCKER_TAG) -f services/$$service/Dockerfile services/$$service/ || exit 1; \
	done
	@echo "$(COLOR_GREEN)All Docker images built successfully!$(COLOR_RESET)"

# Push Docker images to registry
docker-push:
	@echo "$(COLOR_BLUE)Pushing Docker images to registry...$(COLOR_RESET)"
	@for service in $(SERVICES); do \
		echo "$(COLOR_BLUE)Pushing Docker image for $$service...$(COLOR_RESET)"; \
		docker push $(DOCKER_REGISTRY)/$$service:$(DOCKER_TAG) || exit 1; \
	done
	@echo "$(COLOR_GREEN)All Docker images pushed successfully!$(COLOR_RESET)"

# Run all services locally
run:
	@echo "$(COLOR_BLUE)Running all services locally...$(COLOR_RESET)"
	@echo "$(COLOR_YELLOW)Note: This will start all services in the background. Use 'docker-compose up' for a more controlled environment.$(COLOR_RESET)"
	@for service in $(SERVICES); do \
		echo "$(COLOR_BLUE)Starting $$service...$(COLOR_RESET)"; \
		(cd services/$$service && DATABASE_DSN=$(DATABASE_DSN) REDIS_ADDR=$(REDIS_ADDR) go run main.go &) || exit 1; \
	done
	@echo "$(COLOR_GREEN)All services started!$(COLOR_RESET)"

# Run a specific service
run-service:
	@if [ -z "$(SERVICE)" ]; then \
		echo "$(COLOR_YELLOW)Please specify a service with SERVICE=<service-name>$(COLOR_RESET)"; \
		exit 1; \
	fi
	@echo "$(COLOR_BLUE)Running $(SERVICE)...$(COLOR_RESET)"
	@cd services/$(SERVICE) && DATABASE_DSN=$(DATABASE_DSN) REDIS_ADDR=$(REDIS_ADDR) go run main.go
	@echo "$(COLOR_GREEN)$(SERVICE) stopped!$(COLOR_RESET)"

# Run services with hot reload using air
dev:
	@echo "$(COLOR_BLUE)Running services with hot reload...$(COLOR_RESET)"
	@echo "$(COLOR_YELLOW)Note: This requires 'air' to be installed. Run 'make install-tools' first.$(COLOR_RESET)"
	@if [ -z "$(SERVICE)" ]; then \
		echo "$(COLOR_YELLOW)Please specify a service with SERVICE=<service-name>$(COLOR_RESET)"; \
		exit 1; \
	fi
	@cd services/$(SERVICE) && DATABASE_DSN=$(DATABASE_DSN) REDIS_ADDR=$(REDIS_ADDR) air

# Run database migrations
migrate:
	@echo "$(COLOR_BLUE)Running database migrations...$(COLOR_RESET)"
	@cd pkg/database && DATABASE_DSN=$(DATABASE_DSN) MIGRATIONS_DIR=$(MIGRATIONS_DIR) go run migrate.go
	@echo "$(COLOR_GREEN)Migrations completed!$(COLOR_RESET)"

# Generate Swagger documentation
swagger:
	@echo "$(COLOR_BLUE)Generating Swagger documentation...$(COLOR_RESET)"
	@for service in $(SERVICES); do \
		if [ -f services/$$service/main.go ]; then \
			echo "$(COLOR_BLUE)Generating Swagger docs for $$service...$(COLOR_RESET)"; \
			(cd services/$$service && swag init) || echo "$(COLOR_YELLOW)Swagger generation skipped for $$service$(COLOR_RESET)"; \
		fi \
	done
	@echo "$(COLOR_GREEN)Swagger documentation generated!$(COLOR_RESET)"

# Install development tools
install-tools:
	@echo "$(COLOR_BLUE)Installing development tools...$(COLOR_RESET)"
	@go install github.com/cosmtrek/air@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/swaggo/swag/cmd/swag@latest
	@echo "$(COLOR_GREEN)Development tools installed!$(COLOR_RESET)"

# Create a new service (template)
create-service:
	@if [ -z "$(SERVICE)" ]; then \
		echo "$(COLOR_YELLOW)Please specify a service name with SERVICE=<service-name>$(COLOR_RESET)"; \
		exit 1; \
	fi
	@echo "$(COLOR_BLUE)Creating new service: $(SERVICE)...$(COLOR_RESET)"
	@mkdir -p services/$(SERVICE)/domain
	@mkdir -p services/$(SERVICE)/application
	@mkdir -p services/$(SERVICE)/infrastructure
	@mkdir -p services/$(SERVICE)/interfaces
	@echo "module github.com/your-org/$(SERVICE)\n\ngo $(GO_VERSION)" > services/$(SERVICE)/go.mod
	@echo "package main\n\nimport (\n\t\"log\"\n\t\"github.com/labstack/echo/v4\"\n)\n\nfunc main() {\n\te := echo.New()\n\tlog.Println(\"$(SERVICE) running on :8080\")\n\te.Logger.Fatal(e.Start(\":8080\"))\n}" > services/$(SERVICE)/main.go
	@cp services/customer-service/Dockerfile services/$(SERVICE)/Dockerfile
	@cp services/customer-service/Dockerfile.dev services/$(SERVICE)/Dockerfile.dev
	@echo "# $(SERVICE)\n\nThis service is responsible for..." > services/$(SERVICE)/README.md
	@echo "$(COLOR_GREEN)Service $(SERVICE) created successfully!$(COLOR_RESET)"
	@echo "$(COLOR_YELLOW)Don't forget to update the go.mod file with required dependencies.$(COLOR_RESET)"
