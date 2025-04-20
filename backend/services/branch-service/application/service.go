package application

import "../domain"

type BranchService struct {
	repo domain.BranchRepository
}

func NewBranchService(repo domain.BranchRepository) *BranchService {
	return &BranchService{repo: repo}
}

func (s *BranchService) CreateBranch(branch *domain.Branch) error {
	return s.repo.Create(branch)
}

func (s *BranchService) GetBranchByID(id uint) (*domain.Branch, error) {
	return s.repo.GetByID(id)
}

func (s *BranchService) ListBranches() ([]domain.Branch, error) {
	return s.repo.List()
}

func (s *BranchService) UpdateBranch(branch *domain.Branch) error {
	return s.repo.Update(branch)
}

func (s *BranchService) DeleteBranch(id uint) error {
	return s.repo.Delete(id)
}
