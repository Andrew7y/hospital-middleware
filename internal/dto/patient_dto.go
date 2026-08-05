package dto

import (
	"time"
)

type PatientSearchRequest struct {
	FirstName   *string    `form:"first_name" json:"first_name"`
	MiddleName  *string    `form:"middle_name" json:"middle_name"`
	LastName    *string    `form:"last_name" json:"last_name"`
	NationalID  *string    `form:"national_id" json:"national_id"`
	PassportID  *string    `form:"passport_id" json:"passport_id"`
	DateOfBirth *time.Time `form:"date_of_birth" json:"date_of_birth" time_format:"02/01/2006"`
	PhoneNumber *string    `form:"phone_number" json:"phone_number"`
	Email       *string    `form:"email" json:"email"`

	Page     int `form:"page" json:"page" binding:"omitempty,min=1"`
	PageSize int `form:"page_size" json:"page_size" binding:"omitempty,min=1,max=100"`

	SortBy    string `form:"sort_by" json:"sort_by"`
	SortOrder string `form:"sort_order" json:"sort_order"`
}

type PatientSearchResult struct {
	ID           uint    `json:"id"`
	HospitalID   uint    `json:"hospital_id"`
	FirstNameTH  string  `json:"first_name_th"`
	MiddleNameTH string  `json:"middle_name_th"`
	LastNameTH   string  `json:"last_name_th"`
	FirstNameEN  string  `json:"first_name_en"`
	MiddleNameEN string  `json:"middle_name_en"`
	LastNameEN   string  `json:"last_name_en"`
	DateOfBirth  string  `json:"date_of_birth"`
	PatientHN    string  `json:"patient_hn"`
	NationalID   *string `json:"national_id"`
	PassportID   *string `json:"passport_id"`
	PhoneNumber  string  `json:"phone_number"`
	Email        string  `json:"email"`
	Gender       string  `json:"gender"`
}

type PatientSearchResponse struct {
	Patients   []PatientSearchResult `json:"patients"`
	Pagination PaginationMeta        `json:"pagination"`
	Sorting    SortingMeta           `json:"sorting"`
}
