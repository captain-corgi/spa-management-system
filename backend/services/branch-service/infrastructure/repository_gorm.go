package infrastructure

import (
	"github.com/your-org/branch-service/domain"
	"gorm.io/gorm"
)

type GormBranchRepository struct {
	db *gorm.DB
}

func NewGormBranchRepository(db *gorm.DB) *GormBranchRepository {
	return &GormBranchRepository{db: db}
}

func (r *GormBranchRepository) Create(branch *domain.Branch) error {
	return r.db.Create(branch).Error
}

func (r *GormBranchRepository) GetByID(id uint) (*domain.Branch, error) {
	var b domain.Branch
	err := r.db.First(&b, id).Error
	return &b, err
}

func (r *GormBranchRepository) Update(branch *domain.Branch) error {
	return r.db.Save(branch).Error
}

func (r *GormBranchRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Branch{}, id).Error
}

func (r *GormBranchRepository) List() ([]domain.Branch, error) {
	var branches []domain.Branch
	err := r.db.Find(&branches).Error
	return branches, err
}
