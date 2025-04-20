package domain

type Metric struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string
	Value     float64
	Timestamp int64
}

type Report struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string
	CreatedAt int64
}

type Dashboard struct {
	ID    uint   `gorm:"primaryKey"`
	Title string
}
