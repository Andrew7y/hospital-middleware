package normalize

import (
	"hospital-middleware/internal/dto"
	"hospital-middleware/pkg/query"
	"strings"
)

func PatientSorting(sortBy, sortOrder string) query.Sorting {
	sortBy = strings.ToLower(strings.TrimSpace(sortBy))
	sortOrder = strings.ToLower(strings.TrimSpace(sortOrder))

	if sortBy == "" {
		sortBy = "patient_hn"
	}

	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "asc"
	}

	return query.Sorting{
		SortBy:    sortBy,
		SortOrder: sortOrder,
	}
}

func PatientQueryFilter(req dto.PatientSearchRequest) *query.PatientQueryFilter {
	return &query.PatientQueryFilter{
		FirstName:   req.FirstName,
		MiddleName:  req.MiddleName,
		LastName:    req.LastName,
		NationalID:  req.NationalID,
		PassportID:  req.PassportID,
		DateOfBirth: req.DateOfBirth,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
	}
}
