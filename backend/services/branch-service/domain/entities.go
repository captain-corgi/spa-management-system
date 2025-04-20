package domain

type Branch struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string
	Address string
}

type Location struct {
	ID       uint   `gorm:"primaryKey"`
	BranchID uint
	Address  string
}

type Resource struct {
	ID       uint   `gorm:"primaryKey"`
	BranchID uint
	Name     string
	Type     string
}
