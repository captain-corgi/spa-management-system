package domain

type Campaign struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string
	StartDate int64
	EndDate   int64
	Budget    float64
}

type Promotion struct {
	ID         uint   `gorm:"primaryKey"`
	CampaignID uint
	Title      string
	Discount   float64
}

type Engagement struct {
	ID         uint   `gorm:"primaryKey"`
	CampaignID uint
	CustomerID uint
	EngagedAt  int64
}
