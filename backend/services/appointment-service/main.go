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

	repo := infrastructure.NewGormAppointmentRepository(db)
	service := application.NewAppointmentService(repo)

	e := echo.New()
	interfaces.RegisterRoutes(e, service)

	log.Println("Appointment Service running on :8081")
	e.Logger.Fatal(e.Start(":8081"))
}
