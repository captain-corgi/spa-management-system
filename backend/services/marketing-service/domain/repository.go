package domain

type CampaignRepository interface {
	Create(campaign *Campaign) error
	GetByID(id uint) (*Campaign, error)
	Update(campaign *Campaign) error
	Delete(id uint) error
	List() ([]Campaign, error)
}

type PromotionRepository interface {
	Create(promotion *Promotion) error
	ListByCampaign(campaignID uint) ([]Promotion, error)
}

type EngagementRepository interface {
	Create(engagement *Engagement) error
	ListByCampaign(campaignID uint) ([]Engagement, error)
}
