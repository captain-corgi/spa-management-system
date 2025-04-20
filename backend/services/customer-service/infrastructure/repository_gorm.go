package infrastructure

import (
	"gorm.io/gorm"
	"../domain"
)

type GormCustomerRepository struct {
	db *gorm.DB
}

func NewGormCustomerRepository(db *gorm.DB) *GormCustomerRepository {
	return &GormCustomerRepository{db: db}
}

func (r *GormCustomerRepository) Create(customer *domain.Customer) error {
	return r.db.Create(customer).Error
}

func (r *GormCustomerRepository) GetByID(id uint) (*domain.Customer, error) {
	var c domain.Customer
	err := r.db.First(&c, id).Error
	return &c, err
}

func (r *GormCustomerRepository) GetByPhone(phone string) (*domain.Customer, error) {
	var c domain.Customer
	err := r.db.Where("phone = ?", phone).First(&c).Error
	return &c, err
}

func (r *GormCustomerRepository) Update(customer *domain.Customer) error {
	return r.db.Save(customer).Error
}

func (r *GormCustomerRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Customer{}, id).Error
}

func (r *GormCustomerRepository) List() ([]domain.Customer, error) {
	var customers []domain.Customer
	err := r.db.Find(&customers).Error
	return customers, err
}
