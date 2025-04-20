package infrastructure

import (
	"gorm.io/gorm"
	"../domain"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.Customer{},
		&domain.TreatmentRecord{},
		&domain.Purchase{},
		&domain.Debt{},
	)
}
