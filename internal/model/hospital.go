package model

type Hospital struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"not null"`
	APIURL string `gorm:"not null"`
}
