package infrastructure

import (
	"github.com/your-org/staff-service/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.Staff{},
		&domain.Schedule{},
		&domain.PerformanceRecord{},
	)
}
