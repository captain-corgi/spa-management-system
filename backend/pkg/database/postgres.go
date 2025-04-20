package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewPostgresDB creates a new GORM DB instance with connection pooling.
func NewPostgresDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to Postgres: %v", err)
	}
	// Set connection pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB from GORM: %v", err)
	}
	maxOpen := 20
	maxIdle := 10
	maxLifetime := 3600 // seconds
	if v := os.Getenv("DB_MAX_OPEN_CONNS"); v != "" {
		fmt.Sscanf(v, "%d", &maxOpen)
	}
	if v := os.Getenv("DB_MAX_IDLE_CONNS"); v != "" {
		fmt.Sscanf(v, "%d", &maxIdle)
	}
	if v := os.Getenv("DB_CONN_MAX_LIFETIME"); v != "" {
		fmt.Sscanf(v, "%d", &maxLifetime)
	}
	sqlDB.SetMaxOpenConns(maxOpen)
	sqlDB.SetMaxIdleConns(maxIdle)
	sqlDB.SetConnMaxLifetime(time.Duration(maxLifetime) * time.Second)
	return db
}

// WithTransaction runs the given function within a transaction.
func WithTransaction(db *gorm.DB, fn func(tx *gorm.DB) error) error {
	return db.Transaction(fn)
}
