package model

import (
	"time"
)

type Patient struct {
	ID          uint      `gorm:"primaryKey"`
	DateOfBirth time.Time `gorm:"not null;type:date"`
	NationalID  *string   `gorm:"size:20;uniqueIndex"`
	Gender      string
}
