package infrastructure

import (
	"gorm.io/gorm"
	"../domain"
)

type GormInvoiceRepository struct {
	db *gorm.DB
}

func NewGormInvoiceRepository(db *gorm.DB) *GormInvoiceRepository {
	return &GormInvoiceRepository{db: db}
}

func (r *GormInvoiceRepository) Create(invoice *domain.Invoice) error {
	return r.db.Create(invoice).Error
}

func (r *GormInvoiceRepository) GetByID(id uint) (*domain.Invoice, error) {
	var i domain.Invoice
	err := r.db.First(&i, id).Error
	return &i, err
}

func (r *GormInvoiceRepository) Update(invoice *domain.Invoice) error {
	return r.db.Save(invoice).Error
}

func (r *GormInvoiceRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Invoice{}, id).Error
}

func (r *GormInvoiceRepository) List() ([]domain.Invoice, error) {
	var invoices []domain.Invoice
	err := r.db.Find(&invoices).Error
	return invoices, err
}
