package domain

// Appointment represents a spa appointment
 type Appointment struct {
	ID         uint   `gorm:"primaryKey"`
	CustomerID uint
	StaffID    uint
	RoomID     uint
	StartTime  int64
	EndTime    int64
	Status     string
	Notes      string
}

// Room represents a treatment room
 type Room struct {
	ID   uint   `gorm:"primaryKey"`
	Name string
}
