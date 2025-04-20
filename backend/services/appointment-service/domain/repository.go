package domain

type AppointmentRepository interface {
	Create(appointment *Appointment) error
	GetByID(id uint) (*Appointment, error)
	Update(appointment *Appointment) error
	Delete(id uint) error
	List() ([]Appointment, error)
}
