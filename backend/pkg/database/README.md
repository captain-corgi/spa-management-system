# Database Package

This package provides database configuration, connection pooling, migrations, Redis caching, and transaction management for the Spa Management System.

## Structure

- `postgres.go` - PostgreSQL connection and pooling
- `redis.go` - Redis client configuration
- `migrations/` - SQL migration scripts (per domain)
- `.windsurf-guide.md` - Database configuration and guidelines

## Features
- **PostgreSQL**: GORM-based connection, connection pooling, transaction helpers
- **Migrations**: SQL migration scripts for all domains, versioned
- **Redis**: Client setup for caching/session, using go-redis
- **Transaction Management**: Helper functions for DB transactions

## Usage
- Add new domain schemas in `migrations/<domain>_init.sql`
- Run migrations using your preferred tool (e.g. golang-migrate)
- Use `NewPostgresDB()` and `NewRedisClient()` in your services
- See `.windsurf-guide.md` for best practices
