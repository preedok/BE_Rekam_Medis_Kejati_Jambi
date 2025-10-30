package models

import (
	"time"
)

type Appointment struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	PatientID   string    `gorm:"not null" json:"patient_id"`
	PatientName string    `json:"patient_name"`
	DoctorID    string    `gorm:"not null" json:"doctor_id"`
	DoctorName  string    `json:"doctor_name"`
	Poli        string    `json:"poli"`
	Date        string    `gorm:"not null" json:"date"`
	Time        string    `json:"time"`
	Complaint   string    `json:"complaint"`
	Status      string    `gorm:"default:'pending'" json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Appointment) TableName() string {
	return "appointments"
}
