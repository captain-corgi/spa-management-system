# Developer Onboarding Guide

Welcome to the Spa Management System team! This guide will help you set up your environment, understand project structure, and follow our development workflow.

---

## 1. Prerequisites
- Git, Docker, Node.js (LTS), Go (>=1.20), PostgreSQL, Redis
- Access to company GitHub, DockerHub, and CI/CD tools

## 2. Clone & Install
```sh
git clone <repo-url>
cd spa-management-system
```
- Install frontend dependencies: `cd frontend/shell && npm ci`
- Install backend dependencies: `cd backend/services/customer-service && go mod download` (repeat for each service)

## 3. Environment Setup
- Copy `.env.example` files to `.env` and fill in secrets (see `.windsurf-guide.md` for required variables)
- Start PostgreSQL and Redis (Docker Compose or local)

## 4. Running Locally
- **Backend:** `go run main.go` in each service folder
- **Frontend:** `npm run dev` in each app folder
- **Migrations:** `cd backend/pkg/database && make migrate`

## 5. Coding Standards
- Follow DDD, Clean Architecture, and microservice/micro frontend boundaries
- TypeScript for frontend, Go for backend
- Write unit tests for all domain logic

## 6. CI/CD & Deployment
- All changes are tested and deployed via Jenkins and GitHub Actions
- See [../infrastructure/ci-cd/README.md](../infrastructure/ci-cd/README.md) for details

## 7. Support
- For issues, check [troubleshooting.md](./troubleshooting.md) or ask in the team Slack
