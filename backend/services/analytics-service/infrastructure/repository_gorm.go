package infrastructure

import (
	"gorm.io/gorm"
	"../domain"
)

type GormMetricRepository struct {
	db *gorm.DB
}

func NewGormMetricRepository(db *gorm.DB) *GormMetricRepository {
	return &GormMetricRepository{db: db}
}

func (r *GormMetricRepository) Create(metric *domain.Metric) error {
	return r.db.Create(metric).Error
}

func (r *GormMetricRepository) List() ([]domain.Metric, error) {
	var metrics []domain.Metric
	err := r.db.Find(&metrics).Error
	return metrics, err
}
