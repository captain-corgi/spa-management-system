package database

import (
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
)

// RunMigrations runs all SQL migrations in the migrations directory for the given database.
func RunMigrations(db *gorm.DB, migrationsDir string) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+filepath.Clean(migrationsDir),
		"postgres", driver,
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err.Error() != "no change" {
		return err
	}
	log.Println("Migrations applied successfully.")
	return nil
}

// CLI entrypoint for CI/CD
func main() {
	db := NewPostgresDB()
	migrationsDir := os.Getenv("MIGRATIONS_DIR")
	if migrationsDir == "" {
		migrationsDir = "./backend/pkg/database/migrations"
	}
	if err := RunMigrations(db, migrationsDir); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}
