package infrastructure

import (
	"github.com/your-org/marketing-service/domain"
	"gorm.io/gorm"
)

type GormCampaignRepository struct {
	db *gorm.DB
}

func NewGormCampaignRepository(db *gorm.DB) *GormCampaignRepository {
	return &GormCampaignRepository{db: db}
}

func (r *GormCampaignRepository) Create(campaign *domain.Campaign) error {
	return r.db.Create(campaign).Error
}

func (r *GormCampaignRepository) GetByID(id uint) (*domain.Campaign, error) {
	var c domain.Campaign
	err := r.db.First(&c, id).Error
	return &c, err
}

func (r *GormCampaignRepository) Update(campaign *domain.Campaign) error {
	return r.db.Save(campaign).Error
}

func (r *GormCampaignRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Campaign{}, id).Error
}

func (r *GormCampaignRepository) List() ([]domain.Campaign, error) {
	var campaigns []domain.Campaign
	err := r.db.Find(&campaigns).Error
	return campaigns, err
}
