package infrastructure

import (
	"github.com/your-org/marketing-service/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.Campaign{},
		&domain.Promotion{},
		&domain.Engagement{},
	)
}
