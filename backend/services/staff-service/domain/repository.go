package domain

type StaffRepository interface {
	Create(staff *Staff) error
	GetByID(id uint) (*Staff, error)
	Update(staff *Staff) error
	Delete(id uint) error
	List() ([]Staff, error)
}

type ScheduleRepository interface {
	Create(schedule *Schedule) error
	ListByStaff(staffID uint) ([]Schedule, error)
}

type PerformanceRecordRepository interface {
	Create(record *PerformanceRecord) error
	ListByStaff(staffID uint) ([]PerformanceRecord, error)
}
