package models

import (
	"time"
)

type MedicalRecord struct {
	ID            string    `gorm:"primaryKey" json:"id"`
	AppointmentID string    `json:"appointment_id"`
	PatientID     string    `gorm:"not null" json:"patient_id"`
	PatientName   string    `json:"patient_name"`
	Poli          string    `json:"poli"`
	Date          string    `json:"date"`
	Anamnesa      string    `json:"anamnesa"`
	Objective     string    `json:"objective"`
	Diagnosis     string    `json:"diagnosis"`
	Therapy       string    `json:"therapy"`
	Prescription  string    `json:"prescription"`
	NextVisit     string    `json:"next_visit"`
	CreatedBy     string    `json:"created_by"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (MedicalRecord) TableName() string {
	return "medical_records"
}
