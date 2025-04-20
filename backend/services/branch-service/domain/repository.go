package domain

type BranchRepository interface {
	Create(branch *Branch) error
	GetByID(id uint) (*Branch, error)
	Update(branch *Branch) error
	Delete(id uint) error
	List() ([]Branch, error)
}

type ResourceRepository interface {
	Create(resource *Resource) error
	ListByBranch(branchID uint) ([]Resource, error)
}
