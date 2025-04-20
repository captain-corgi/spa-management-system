# Database Migrations

This directory contains versioned SQL migration scripts for all domains in the Spa Management System.

## Structure

- `customer/` - Customer domain migrations
- `appointment/` - Appointment domain migrations
- `staff/` - Staff domain migrations
- `branch/` - Branch domain migrations
- `marketing/` - Marketing domain migrations
- `finance/` - Finance domain migrations
- `analytics/` - Analytics domain migrations

## Usage
- Place each domain's migration scripts in its respective subfolder.
- Use a migration tool (e.g., golang-migrate) to apply migrations.
- Use semantic versioning for migration files (e.g., `V1__init.sql`, `V2__add_column.sql`).
