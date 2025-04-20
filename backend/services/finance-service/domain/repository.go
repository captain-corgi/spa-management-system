package domain

type InvoiceRepository interface {
	Create(invoice *Invoice) error
	GetByID(id uint) (*Invoice, error)
	Update(invoice *Invoice) error
	Delete(id uint) error
	List() ([]Invoice, error)
}

type PaymentRepository interface {
	Create(payment *Payment) error
	ListByInvoice(invoiceID uint) ([]Payment, error)
}

type RevenueRepository interface {
	Create(revenue *Revenue) error
	List() ([]Revenue, error)
}

type ExpenseRepository interface {
	Create(expense *Expense) error
	List() ([]Expense, error)
}
