package domain

type Staff struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string
	Role      string
	Phone     string
	Email     string
	CreatedAt int64
	UpdatedAt int64
}

type Schedule struct {
	ID      uint   `gorm:"primaryKey"`
	StaffID uint
	Day     string
	Start   string
	End     string
}

type PerformanceRecord struct {
	ID      uint   `gorm:"primaryKey"`
	StaffID uint
	Score   float64
	Date    int64
}
