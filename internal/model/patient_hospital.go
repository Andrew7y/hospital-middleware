package model

type PatientHospital struct {
	PatientID    uint    `gorm:"primaryKey;autoIncrement:false"`
	HospitalID   uint    `gorm:"primaryKey;autoIncrement:false;uniqueIndex:uq_hospital_hn"`
	FirstNameTH  string  `gorm:"size:100;index:idx_patient_name_th"`
	LastNameTH   string  `gorm:"size:100;index:idx_patient_name_th"`
	MiddleNameTH string  `gorm:"size:100;index:idx_patient_name_th"`
	FirstNameEN  string  `gorm:"size:100;index:idx_patient_name_en"`
	MiddleNameEN string  `gorm:"size:100;index:idx_patient_name_en"`
	LastNameEN   string  `gorm:"size:100;index:idx_patient_name_en"`
	PassportID   *string `gorm:"size:20;uniqueIndex"`
	PhoneNumber  string  `gorm:"size:20;index"`
	Email        string  `gorm:"size:100;index"`
	PatientHN    string  `gorm:"size:50;not null;uniqueIndex:uq_hospital_hn"`

	Patient  Patient  `gorm:"foreignKey:PatientID"`
	Hospital Hospital `gorm:"foreignKey:HospitalID"`
}
