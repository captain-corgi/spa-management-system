package domain

// Customer represents a spa customer
 type Customer struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string
	Phone     string `gorm:"unique"`
	Email     string
	Address   string
	CreatedAt int64
	UpdatedAt int64
}

// TreatmentRecord represents a customer's treatment history
 type TreatmentRecord struct {
	ID          uint   `gorm:"primaryKey"`
	CustomerID  uint
	StaffID     uint
	Treatment   string
	Date        int64
	Notes       string
}

// Purchase represents a purchase made by a customer
 type Purchase struct {
	ID          uint   `gorm:"primaryKey"`
	CustomerID  uint
	Item        string
	Amount      float64
	PurchasedAt int64
}

// Debt represents a customer's debt record
 type Debt struct {
	ID          uint   `gorm:"primaryKey"`
	CustomerID  uint
	Amount      float64
	Reason      string
	CreatedAt   int64
}
