package models

import (
	"time"

	"github.com/lib/pq"
)

type Doctor struct {
	ID             string         `gorm:"primaryKey" json:"id"`
	Name           string         `gorm:"not null" json:"name"`
	Email          string         `gorm:"unique;not null" json:"email"`
	Password       string         `gorm:"not null" json:"-"`
	Schedule       string         `json:"schedule"`
	AvailableTimes pq.StringArray `gorm:"type:text[]" json:"available_times"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}

func (Doctor) TableName() string {
	return "doctors"
}
