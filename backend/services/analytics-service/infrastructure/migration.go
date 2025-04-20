package infrastructure

import (
	"github.com/your-org/analytics-service/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.Metric{},
		&domain.Report{},
		&domain.Dashboard{},
	)
}
