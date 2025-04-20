package infrastructure

import (
	"github.com/your-org/staff-service/domain"
	"gorm.io/gorm"
)

type GormStaffRepository struct {
	db *gorm.DB
}

func NewGormStaffRepository(db *gorm.DB) *GormStaffRepository {
	return &GormStaffRepository{db: db}
}

func (r *GormStaffRepository) Create(staff *domain.Staff) error {
	return r.db.Create(staff).Error
}

func (r *GormStaffRepository) GetByID(id uint) (*domain.Staff, error) {
	var s domain.Staff
	err := r.db.First(&s, id).Error
	return &s, err
}

func (r *GormStaffRepository) Update(staff *domain.Staff) error {
	return r.db.Save(staff).Error
}

func (r *GormStaffRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Staff{}, id).Error
}

func (r *GormStaffRepository) List() ([]domain.Staff, error) {
	var staff []domain.Staff
	err := r.db.Find(&staff).Error
	return staff, err
}
