package application

import "../domain"

type MarketingService struct {
	repo domain.CampaignRepository
}

func NewMarketingService(repo domain.CampaignRepository) *MarketingService {
	return &MarketingService{repo: repo}
}

func (s *MarketingService) CreateCampaign(c *domain.Campaign) error {
	return s.repo.Create(c)
}

func (s *MarketingService) GetCampaignByID(id uint) (*domain.Campaign, error) {
	return s.repo.GetByID(id)
}

func (s *MarketingService) ListCampaigns() ([]domain.Campaign, error) {
	return s.repo.List()
}

func (s *MarketingService) UpdateCampaign(c *domain.Campaign) error {
	return s.repo.Update(c)
}

func (s *MarketingService) DeleteCampaign(id uint) error {
	return s.repo.Delete(id)
}
