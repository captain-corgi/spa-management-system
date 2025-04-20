# Spa Management System Documentation

Welcome to the documentation hub for the Spa Management System. Here you'll find comprehensive guides on architecture, APIs, setup, development, and testing.

---

## 📐 System Architecture
- **Overview:**
  - Microservices backend (Go, GORM, Echo)
  - Micro frontend (React, TypeScript, Vite, Module Federation)
  - PostgreSQL, Redis, Kubernetes, Docker
- **Diagram:**
  - ![System Architecture Diagram](./system-architecture.png) <!-- Replace with actual diagram file -->
- **See:** [integration-strategy.md](./integration-strategy.md)

## 📚 API Documentation
- **Backend APIs:** OpenAPI/Swagger per service (see `/backend/services/*/docs`)
- **API Gateway:** [api-gateway/README.md](../backend/api-gateway/README.md)
- **Data Contracts:** `/contracts/` (shared schemas)

## 🚀 Setup & Deployment Guides
- **CI/CD:** [../infrastructure/ci-cd/README.md](../infrastructure/ci-cd/README.md)
- **Local Development:** [../infrastructure/ci-cd/jenkins/README.md](../infrastructure/ci-cd/jenkins/README.md)
- **Kubernetes:** [../infrastructure/k8s/README.md](../infrastructure/k8s/README.md)

## 🛠️ Development Workflows
- **Backend:** DDD, Clean Architecture, Go modules, GORM, unit tests
- **Frontend:** Micro frontend, TypeScript, Vite, shared state, unit/E2E tests
- **Environment Setup:** See setup guides above

## ✅ Testing Strategy
- **Unit Tests:** Domain logic for backend (Go) and frontend (JS/TS)
- **Integration Tests:** Service-to-service, API gateway, database, message broker
- **End-to-End Tests:** Full user flows, Cypress/Playwright
- **See:** [integration-strategy.md](./integration-strategy.md)

---

For detailed guides, see the referenced files above. For questions or onboarding, contact the project maintainer.
