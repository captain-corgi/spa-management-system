package domain

type Invoice struct {
	ID        uint   `gorm:"primaryKey"`
	Number    string
	CustomerID uint
	Amount    float64
	IssuedAt  int64
}

type Payment struct {
	ID        uint   `gorm:"primaryKey"`
	InvoiceID uint
	Amount    float64
	PaidAt    int64
}

type Revenue struct {
	ID     uint   `gorm:"primaryKey"`
	Month  string
	Amount float64
}

type Expense struct {
	ID     uint   `gorm:"primaryKey"`
	Type   string
	Amount float64
	Date   int64
	Approved bool
}
