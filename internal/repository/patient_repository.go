package repository

import (
	"context"
	"gorm.io/gorm"
	"hospital-middleware/pkg/query"
)

type PatientRepository interface {
	SearchPatient(
		c context.Context,
		hospitalID uint,
		filter *query.PatientQueryFilter,
		pagination query.Pagination,
		sorting query.Sorting,
	) (*query.PatientSearchPage, error)
}

type patientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) PatientRepository {
	return &patientRepository{db: db}
}

func (r *patientRepository) SearchPatient(
	c context.Context,
	hospitalID uint,
	filter *query.PatientQueryFilter,
	pagination query.Pagination,
	sorting query.Sorting,
) (*query.PatientSearchPage, error) {
	db := r.db.WithContext(c).
		Table("patient_hospitals AS ph").
		Joins("INNER JOIN patients AS p ON ph.patient_id = p.id").
		Where("ph.hospital_id = ?", hospitalID)

	db = query.ApplyPatientQueryFilter(db, filter)

	var totalItems int64
	if err := db.Count(&totalItems).Error; err != nil {
		return nil, err
	}

	sortColumn, sortOrder := query.ResolvePatientSortOrder(sorting)

	var rows []query.PatientQueryRow

	err := db.
		Select(
			`p.id AS patient_id,
            ph.hospital_id,
			ph.first_name_th,
			ph.middle_name_th,
			ph.last_name_th,
			ph.first_name_en,
			ph.middle_name_en,
			ph.last_name_en,
			p.date_of_birth,
			ph.patient_hn,
			p.national_id,
			ph.passport_id,
			ph.phone_number,
			ph.email,
			p.gender`).
		Order(sortColumn + " " + sortOrder).
		Limit(pagination.PageSize).
		Offset(pagination.Offset()).
		Scan(&rows).
		Error

	if err != nil {
		return nil, err
	}

	return &query.PatientSearchPage{
		Items:      rows,
		TotalItems: totalItems,
	}, nil
}
