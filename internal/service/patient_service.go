package service

import (
	"context"
	"hospital-middleware/internal/dto"
	"hospital-middleware/internal/repository"
	"hospital-middleware/pkg/normalize"
)

type PatientService interface {
	SearchPatient(
		c context.Context,
		hospitalID uint,
		req dto.PatientSearchRequest,
	) (*dto.PatientSearchResponse, error)
}

type patientService struct {
	patientRepo repository.PatientRepository
}

func NewPatientService(repo repository.PatientRepository) PatientService {
	return &patientService{
		patientRepo: repo,
	}
}

func (s *patientService) SearchPatient(
	c context.Context,
	hospitalID uint,
	req dto.PatientSearchRequest,
) (*dto.PatientSearchResponse, error) {
	pagination := normalize.PaginationValue(req.Page, req.PageSize)
	sorting := normalize.PatientSorting(req.SortBy, req.SortOrder)
	filter := normalize.PatientQueryFilter(req)

	result, err := s.patientRepo.SearchPatient(
		c,
		hospitalID,
		filter,
		pagination,
		sorting,
	)
	if err != nil {
		return nil, err
	}

	items := make([]dto.PatientSearchResult, len(result.Items))

	for i, row := range result.Items {
		items[i] = dto.PatientSearchResult{
			ID:           row.PatientID,
			HospitalID:   row.HospitalID,
			FirstNameTH:  row.FirstNameTH,
			MiddleNameTH: row.MiddleNameTH,
			LastNameTH:   row.LastNameTH,
			FirstNameEN:  row.FirstNameEN,
			MiddleNameEN: row.MiddleNameEN,
			LastNameEN:   row.LastNameEN,
			DateOfBirth:  row.DateOfBirth.Format("02/01/2006"),
			PatientHN:    row.PatientHN,
			NationalID:   row.NationalID,
			PassportID:   row.PassportID,
			PhoneNumber:  row.PhoneNumber,
			Email:        row.Email,
			Gender:       row.Gender,
		}
	}

	totalPages := 0
	if result.TotalItems > 0 {
		totalPages = int(
			(result.TotalItems + int64(pagination.PageSize) - 1) /
				int64(pagination.PageSize),
		)
	}

	return &dto.PatientSearchResponse{
		Patients: items,
		Pagination: dto.PaginationMeta{
			Page:       pagination.Page,
			PageSize:   pagination.PageSize,
			TotalItems: result.TotalItems,
			TotalPages: totalPages,
		},
		Sorting: dto.SortingMeta{
			SortBy:    sorting.SortBy,
			SortOrder: sorting.SortOrder,
		},
	}, nil

}
