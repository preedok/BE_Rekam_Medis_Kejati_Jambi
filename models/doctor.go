package models

import (
	"time"
)

type Patient struct {
	ID             string    `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"not null" json:"name"`
	Email          string    `gorm:"unique;not null" json:"email"`
	Password       string    `gorm:"not null" json:"-"`
	Phone          string    `json:"phone"`
	Address        string    `json:"address"`
	DateOfBirth    string    `json:"date_of_birth"`
	Gender         string    `json:"gender"`
	BloodType      string    `json:"blood_type"`
	RegisteredDate time.Time `json:"registered_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (Patient) TableName() string {
	return "patients"
}
