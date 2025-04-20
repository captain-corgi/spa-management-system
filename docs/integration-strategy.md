# Integration Strategy

This document describes the integration patterns and practices for the Spa Management System, covering service communication, frontend integration, testing, data contracts, and API gateway configuration.

---

## 1. Service Communication

- **REST APIs:**
  - All backend services expose RESTful endpoints for synchronous operations.
  - OpenAPI/Swagger is used for API documentation and contract sharing.
- **Message Broker:**
  - Asynchronous events are published to a message broker (e.g., Kafka, RabbitMQ) for decoupled communication between services.
  - Used for notifications, audit logs, and cross-service workflows.
- **Shared Data Contracts:**
  - All request/response and event schemas are defined in a central `contracts/` repository or package.
  - Versioning is enforced to prevent breaking changes.

## 2. Frontend Integration

- **Micro Frontend Architecture:**
  - Each domain (customer, appointment, staff, etc.) is a separate React + TypeScript app, loaded via Module Federation.
  - The shell app orchestrates loading and routing of micro frontends.
- **Shared State Management:**
  - Global state (auth, user profile, theme) is managed via a shared state library (e.g., Redux, Zustand) and exposed to all micro frontends.
- **Authentication Token Sharing:**
  - Auth tokens are securely shared between shell and micro frontends via browser storage and context providers.
  - All API requests use the same token for SSO experience.

## 3. Testing Strategy

- **Unit Tests:**
  - All domain logic (backend and frontend) is covered by unit tests (Go: testify, JS: jest).
- **Integration Tests:**
  - Service-to-service and API gateway interactions are tested using integration test suites.
  - Database and message broker are included in test containers.
- **End-to-End Tests:**
  - Critical user flows are covered by E2E tests (e.g., Cypress, Playwright).
  - Run in CI for every PR and nightly on main.

## 4. Data Contracts

- **API Schemas:**
  - All REST APIs are defined with OpenAPI specs and validated at runtime.
- **Event Schemas:**
  - All published events are versioned and validated against shared schemas (e.g., using Avro, JSON Schema).
- **Backward Compatibility:**
  - All changes to contracts are reviewed for compatibility and versioned.

## 5. API Gateway Configuration

- **API Gateway:**
  - All external and internal API traffic flows through the API gateway (e.g., Kong, NGINX, custom Go gateway).
  - Handles routing, authentication, rate limiting, and observability.
- **Configuration:**
  - Each service registers its routes and OpenAPI spec with the gateway.
  - Auth middleware verifies tokens and injects user context.
  - Custom plugins can be added for logging, metrics, and security.

---

For more details, see:
- `/backend/api-gateway/` for gateway implementation and configuration
- `/contracts/` or shared package for data contracts
- Service and frontend READMEs for integration and testing instructions
