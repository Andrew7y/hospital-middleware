package model

type Staff struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"size:120;not null;uniqueIndex"`
	Password string `gorm:"not null"`

	HospitalID uint     `gorm:"index"`
	Hospital   Hospital `gorm:"foreignKey:HospitalID"`
}
