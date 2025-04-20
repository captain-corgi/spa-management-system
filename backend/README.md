# Backend

This directory contains all backend-related code, including microservices, shared packages, and the API gateway. The backend follows Domain-Driven Design (DDD), Clean Architecture, and Go best practices to create a scalable, maintainable, and modular system.

## Overview

The Spa Management System backend is built as a collection of microservices, each responsible for a specific business domain. This architecture allows for independent development, deployment, and scaling of services. The system uses:

- **Go (Golang)** as the primary programming language
- **GORM** for object-relational mapping
- **Echo Framework** for RESTful API development
- **PostgreSQL** for persistent storage
- **Redis** for caching and session management
- **Message Broker** for asynchronous communication between services

## Project Structure

```text
backend/
├── api-gateway/         # API Gateway for routing and cross-cutting concerns
├── pkg/                 # Shared packages used across services
│   ├── auth/            # Authentication utilities
│   ├── common/          # Common utilities (logging, errors, etc.)
│   └── database/        # Database utilities and migrations
├── services/            # Microservices
│   ├── customer-service/    # Customer management service
│   ├── appointment-service/ # Appointment management service
│   ├── staff-service/       # Staff management service
│   ├── branch-service/      # Branch management service
│   ├── marketing-service/   # Marketing management service
│   ├── finance-service/     # Finance management service
│   └── analytics-service/   # Analytics service
└── Makefile             # Common operations for all services
```

## Makefile Usage

The Makefile provides common operations for all backend services. Here are some usage examples:

### Building Services

```bash
# Build all services
make build

# Build a specific service
make build-service SERVICE=customer-service
```

### Running Tests

```bash
# Run tests for all services
make test

# Run tests for a specific service
make test-service SERVICE=appointment-service
```

### Development

```bash
# Run a service with hot reload
make dev SERVICE=customer-service

# Install development tools (air, golangci-lint, swag)
make install-tools
```

### Database Operations

```bash
# Run database migrations
make migrate

# Run migrations with custom database connection
make migrate DATABASE_DSN="postgres://user:pass@host:port/dbname?sslmode=disable"
```

### Docker Operations

```bash
# Build Docker images for all services
make docker-build

# Push Docker images to registry
make docker-push DOCKER_REGISTRY=your-registry DOCKER_TAG=v1.0.0
```

### Creating a New Service

```bash
# Create a new service with standard folder structure
make create-service SERVICE=new-service-name
```

### Other Commands

```bash
# Generate Swagger documentation
make swagger

# Run linter on all services
make lint

# Show help information
make help
```

## Environment Variables

The following environment variables can be used with the Makefile:

| Variable | Description | Default |
|----------|-------------|----------|
| DATABASE_DSN | Database connection string | postgres://postgres:postgres@localhost:5432/spa?sslmode=disable |
| REDIS_ADDR | Redis address | localhost:6379 |
| SERVICE | Specific service to operate on | - |
| ENV | Environment (dev, sit, uat, nft, prd) | dev |
| DOCKER_REGISTRY | Docker registry | spa-management |
| DOCKER_TAG | Docker tag | latest |

## Component Details

### API Gateway

The API Gateway (`api-gateway/`) serves as the entry point for all client requests and provides the following functionality:

- **Request Routing**: Routes requests to appropriate microservices
- **Authentication**: Verifies JWT tokens and handles authorization
- **Rate Limiting**: Prevents abuse by limiting request rates
- **Request Validation**: Validates incoming requests
- **Response Caching**: Caches responses for improved performance
- **API Documentation**: Provides Swagger documentation for all endpoints
- **Circuit Breaking**: Prevents cascading failures
- **CORS Configuration**: Handles cross-origin requests

### Shared Packages (pkg/)

The `pkg/` directory contains shared code used across multiple services:

#### Authentication (auth/)

- JWT token generation and validation
- Role-based access control
- Authentication middleware

#### Common Utilities (common/)

- Error handling and custom error types
- Logging utilities
- Message broker integration
- Common data structures and constants

#### Database Utilities (database/)

- Database connection management
- Migration tools and scripts
- PostgreSQL integration
- Redis client configuration
- Transaction management

### Microservices (services/)

Each service in the `services/` directory is responsible for a specific business domain:

#### Customer Service

Manages customer data and interactions:

- Customer profiles and personal information
- Treatment history and records
- Purchase history
- Debt tracking and management

#### Appointment Service

Handles all appointment-related functionality:

- Appointment scheduling and booking
- Calendar management
- Room and resource allocation
- Appointment status tracking and notifications

#### Staff Service

Manages staff-related operations:

- Staff profiles and credentials
- Scheduling and availability
- Performance tracking
- Skill and certification management

#### Branch Service

Manages spa locations and facilities:

- Branch information and details
- Location management
- Resource inventory and allocation
- Branch-specific settings

#### Marketing Service

Handles marketing and promotional activities:

- Campaign management
- Promotion creation and tracking
- Customer engagement analytics
- Marketing performance metrics

#### Finance Service

Manages financial operations:

- Invoice generation and management
- Payment processing and tracking
- Revenue reporting
- Expense management

#### Analytics Service

Provides business intelligence and reporting:

- Data collection and aggregation
- Report generation
- Dashboard metrics
- Performance analytics

## Service Architecture

Each service follows Clean Architecture with the following layers:

### Domain Layer

The core of the application containing:

- Business entities and value objects
- Repository interfaces defining data access contracts
- Domain services for complex business logic
- Domain events and event handlers

### Application Layer

Orchestrates the domain objects to perform tasks:

- Service implementations coordinating domain operations
- Use case implementations
- Command and query handlers
- Application events and event handlers

### Infrastructure Layer

Provides technical capabilities and adapters:

- Repository implementations (using GORM)
- Database migrations and schema management
- External service integrations
- Message broker implementations
- Caching mechanisms

### Interfaces Layer

Handles communication with the outside world:

- HTTP handlers and controllers (using Echo)
- Request/response DTOs and validation
- Middleware configuration
- API route definitions
- API documentation

## Communication Patterns

Services communicate using the following patterns:

1. **Synchronous Communication**: REST APIs for direct service-to-service communication
2. **Asynchronous Communication**: Message broker for event-driven communication
3. **Database Integration**: Each service has its own database schema to maintain independence
