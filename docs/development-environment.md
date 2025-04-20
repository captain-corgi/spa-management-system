# Local Development Environment Setup

This guide describes how to set up and run the Spa Management System locally using Docker Compose, with hot reloading, sample data, and environment variable documentation.

---

## 1. Prerequisites
- Docker & Docker Compose
- Node.js (for frontend hot reload)
- Go (for backend hot reload)

## 2. Quick Start
```sh
docker-compose up --build
```
- Starts PostgreSQL, Redis, backend (with hot reload), and frontend (with hot reload)
- Sample data is loaded into PostgreSQL automatically

## 3. Hot Reloading
- **Backend:** Uses [air](https://github.com/cosmtrek/air) for Go hot reload (see `Dockerfile.dev`)
- **Frontend:** Uses Vite/React hot reload (see `Dockerfile.dev`)
- Code changes are reflected instantly in running containers

## 4. Environment Variables
- Copy `.env.example` to `.env` and fill in values as needed
- Main variables:
  - `DATABASE_DSN` (PostgreSQL connection string)
  - `REDIS_ADDR` (Redis address)
  - `VITE_API_URL` (Frontend API base URL)
  - See `.env.example` for all options

## 5. Sample Data
- SQL scripts in `infrastructure/docker/postgres-init/` are loaded at container startup
- Add more sample data as needed for development/testing

## 6. Extending Compose
- To add more services/apps, duplicate the service definition in `docker-compose.yml` and adjust context/ports

---

For troubleshooting, see [troubleshooting.md](./troubleshooting.md).
For onboarding, see [onboarding.md](./onboarding.md).
