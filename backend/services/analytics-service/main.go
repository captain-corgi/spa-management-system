package main

import (
	"log"
	"os"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"./domain"
	"./application"
	"./infrastructure"
	"./interfaces"
)

func main() {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := infrastructure.Migrate(db); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	repo := infrastructure.NewGormMetricRepository(db)
	service := application.NewAnalyticsService(repo)

	e := echo.New()
	interfaces.RegisterRoutes(e, service)

	log.Println("Analytics Service running on :8086")
	e.Logger.Fatal(e.Start(":8086"))
}
