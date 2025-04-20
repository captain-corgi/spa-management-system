package infrastructure

import (
	"github.com/your-org/branch-service/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.Branch{},
		&domain.Location{},
		&domain.Resource{},
	)
}
