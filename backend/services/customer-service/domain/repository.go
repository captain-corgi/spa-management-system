package domain

// CustomerRepository defines the interface for customer persistence
 type CustomerRepository interface {
	Create(customer *Customer) error
	GetByID(id uint) (*Customer, error)
	GetByPhone(phone string) (*Customer, error)
	Update(customer *Customer) error
	Delete(id uint) error
	List() ([]Customer, error)
}

// TreatmentRecordRepository defines the interface for treatment record persistence
 type TreatmentRecordRepository interface {
	Create(record *TreatmentRecord) error
	ListByCustomer(customerID uint) ([]TreatmentRecord, error)
}

// PurchaseRepository defines the interface for purchase persistence
 type PurchaseRepository interface {
	Create(purchase *Purchase) error
	ListByCustomer(customerID uint) ([]Purchase, error)
}

// DebtRepository defines the interface for debt persistence
 type DebtRepository interface {
	Create(debt *Debt) error
	ListByCustomer(customerID uint) ([]Debt, error)
}
