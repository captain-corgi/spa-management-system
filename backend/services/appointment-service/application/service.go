package application

import "../domain"

type AppointmentService struct {
	repo domain.AppointmentRepository
}

func NewAppointmentService(repo domain.AppointmentRepository) *AppointmentService {
	return &AppointmentService{repo: repo}
}

func (s *AppointmentService) CreateAppointment(a *domain.Appointment) error {
	return s.repo.Create(a)
}

func (s *AppointmentService) GetAppointmentByID(id uint) (*domain.Appointment, error) {
	return s.repo.GetByID(id)
}

func (s *AppointmentService) ListAppointments() ([]domain.Appointment, error) {
	return s.repo.List()
}

func (s *AppointmentService) UpdateAppointment(a *domain.Appointment) error {
	return s.repo.Update(a)
}

func (s *AppointmentService) DeleteAppointment(id uint) error {
	return s.repo.Delete(id)
}
