package infrastructure

import (
	"github.com/your-org/appointment-service/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.Appointment{},
		&domain.Room{},
	)
}
