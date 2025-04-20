package infrastructure

import (
	"github.com/your-org/finance-service/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.Invoice{},
		&domain.Payment{},
		&domain.Revenue{},
		&domain.Expense{},
	)
}
