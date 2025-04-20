package infrastructure

import (
	"gorm.io/gorm"
	"../domain"
)

type GormAppointmentRepository struct {
	db *gorm.DB
}

func NewGormAppointmentRepository(db *gorm.DB) *GormAppointmentRepository {
	return &GormAppointmentRepository{db: db}
}

func (r *GormAppointmentRepository) Create(appointment *domain.Appointment) error {
	return r.db.Create(appointment).Error
}

func (r *GormAppointmentRepository) GetByID(id uint) (*domain.Appointment, error) {
	var a domain.Appointment
	err := r.db.First(&a, id).Error
	return &a, err
}

func (r *GormAppointmentRepository) Update(appointment *domain.Appointment) error {
	return r.db.Save(appointment).Error
}

func (r *GormAppointmentRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Appointment{}, id).Error
}

func (r *GormAppointmentRepository) List() ([]domain.Appointment, error) {
	var appointments []domain.Appointment
	err := r.db.Find(&appointments).Error
	return appointments, err
}
