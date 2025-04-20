package application

import "../domain"

type CustomerService struct {
	repo domain.CustomerRepository
}

func NewCustomerService(repo domain.CustomerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

func (s *CustomerService) CreateCustomer(c *domain.Customer) error {
	return s.repo.Create(c)
}

func (s *CustomerService) GetCustomerByID(id uint) (*domain.Customer, error) {
	return s.repo.GetByID(id)
}

func (s *CustomerService) ListCustomers() ([]domain.Customer, error) {
	return s.repo.List()
}

func (s *CustomerService) UpdateCustomer(c *domain.Customer) error {
	return s.repo.Update(c)
}

func (s *CustomerService) DeleteCustomer(id uint) error {
	return s.repo.Delete(id)
}
