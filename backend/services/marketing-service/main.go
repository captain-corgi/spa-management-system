package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/your-org/marketing-service/application"
	"github.com/your-org/marketing-service/infrastructure"
	"github.com/your-org/marketing-service/interfaces"
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

	repo := infrastructure.NewGormCampaignRepository(db)
	service := application.NewMarketingService(repo)

	e := echo.New()
	interfaces.RegisterRoutes(e, service)

	log.Println("Marketing Service running on :8084")
	e.Logger.Fatal(e.Start(":8084"))
}
