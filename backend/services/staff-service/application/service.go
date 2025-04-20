package application

import "../domain"

type StaffService struct {
	repo domain.StaffRepository
}

func NewStaffService(repo domain.StaffRepository) *StaffService {
	return &StaffService{repo: repo}
}

func (s *StaffService) CreateStaff(staff *domain.Staff) error {
	return s.repo.Create(staff)
}

func (s *StaffService) GetStaffByID(id uint) (*domain.Staff, error) {
	return s.repo.GetByID(id)
}

func (s *StaffService) ListStaff() ([]domain.Staff, error) {
	return s.repo.List()
}

func (s *StaffService) UpdateStaff(staff *domain.Staff) error {
	return s.repo.Update(staff)
}

func (s *StaffService) DeleteStaff(id uint) error {
	return s.repo.Delete(id)
}
