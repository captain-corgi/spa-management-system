package infrastructure

import (
	"gorm.io/gorm"
	"../domain"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.Invoice{},
		&domain.Payment{},
		&domain.Revenue{},
		&domain.Expense{},
	)
}
