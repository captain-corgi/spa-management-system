package infrastructure

import (
	"github.com/your-org/customer-service/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.Customer{},
		&domain.TreatmentRecord{},
		&domain.Purchase{},
		&domain.Debt{},
	)
}
