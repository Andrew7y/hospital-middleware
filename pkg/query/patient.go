package query

import (
	"gorm.io/gorm"
	"strings"
	"time"
)

type PatientQueryRow struct {
	PatientID    uint
	HospitalID   uint
	FirstNameTH  string
	MiddleNameTH string
	LastNameTH   string
	FirstNameEN  string
	MiddleNameEN string
	LastNameEN   string
	DateOfBirth  time.Time
	PatientHN    string
	NationalID   *string
	PassportID   *string
	PhoneNumber  string
	Email        string
	Gender       string
}

type PatientQueryFilter struct {
	FirstName   *string
	MiddleName  *string
	LastName    *string
	NationalID  *string
	PassportID  *string
	DateOfBirth *time.Time
	PhoneNumber *string
	Email       *string
}

type PatientSearchPage struct {
	Items      []PatientQueryRow
	TotalItems int64
}

func ApplyPatientQueryFilter(
	db *gorm.DB,
	filter *PatientQueryFilter,
) *gorm.DB {
	if filter == nil {
		return db
	}

	if filter.FirstName != nil && *filter.FirstName != "" {
		keyword := "%" + *filter.FirstName + "%"
		db = db.Where(
			"ph.first_name_th ILIKE ? "+
				"OR ph.first_name_en ILIKE ?",
			keyword,
			keyword,
		)
	}

	if filter.MiddleName != nil && *filter.MiddleName != "" {
		keyword := "%" + *filter.MiddleName + "%"
		db = db.Where(
			"ph.middle_name_th ILIKE ? "+
				"OR ph.middle_name_en ILIKE ?",
			keyword,
			keyword,
		)
	}

	if filter.LastName != nil && *filter.LastName != "" {
		keyword := "%" + *filter.LastName + "%"
		db = db.Where(
			"ph.last_name_th ILIKE ? "+
				"OR ph.last_name_en ILIKE ?",
			keyword,
			keyword,
		)
	}

	if filter.NationalID != nil && *filter.NationalID != "" {
		db = db.Where("p.national_id = ?", *filter.NationalID)
	}

	if filter.PassportID != nil && *filter.PassportID != "" {
		db = db.Where("p.passport_id = ?", *filter.PassportID)
	}

	if filter.DateOfBirth != nil {
		dobStr := filter.DateOfBirth.Format("2006-01-02")
		db = db.Where("p.date_of_birth = ?", dobStr)
	}

	if filter.PhoneNumber != nil && *filter.PhoneNumber != "" {
		db = db.Where("ph.phone_number = ?", *filter.PhoneNumber)
	}

	if filter.Email != nil && *filter.Email != "" {
		lowerEmail := strings.ToLower(*filter.Email)
		db = db.Where("ph.email = ?", lowerEmail)
	}

	return db
}

var patientSortColumns = map[string]string{
	"id":            "p.id",
	"hospital_id":   "ph.hospital_id",
	"first_name_th": "ph.first_name_th",
	"first_name_en": "ph.first_name_en",
	"last_name_th":  "ph.last_name_th",
	"last_name_en":  "ph.last_name_en",
	"date_of_birth": "p.date_of_birth",
	"patient_hn":    "ph.patient_hn",
	"national_id":   "p.national_id",
	"passport_id":   "ph.passport_id",
	"phone_number":  "ph.phone_number",
	"email":         "ph.email",
	"gender":        "p.gender",
}

func ResolvePatientSortOrder(sorting Sorting) (string, string) {
	column, exists := patientSortColumns[sorting.SortBy]
	if !exists {
		column = "ph.patient_hn"
	}

	order := sorting.SortOrder
	if order != "asc" && order != "desc" {
		order = "asc"
	}

	return column, order
}
