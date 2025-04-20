package application

import "../domain"

type FinanceService struct {
	repo domain.InvoiceRepository
}

func NewFinanceService(repo domain.InvoiceRepository) *FinanceService {
	return &FinanceService{repo: repo}
}

func (s *FinanceService) CreateInvoice(invoice *domain.Invoice) error {
	return s.repo.Create(invoice)
}

func (s *FinanceService) GetInvoiceByID(id uint) (*domain.Invoice, error) {
	return s.repo.GetByID(id)
}

func (s *FinanceService) ListInvoices() ([]domain.Invoice, error) {
	return s.repo.List()
}

func (s *FinanceService) UpdateInvoice(invoice *domain.Invoice) error {
	return s.repo.Update(invoice)
}

func (s *FinanceService) DeleteInvoice(id uint) error {
	return s.repo.Delete(id)
}
