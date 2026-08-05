package tests

import (
	"context"
	"github.com/stretchr/testify/mock"
	"hospital-middleware/internal/dto"
)

type MockPatientService struct {
	mock.Mock
}

func (m *MockPatientService) SearchPatient(
	c context.Context,
	hospitalID uint,
	req dto.PatientSearchRequest,
) (*dto.PatientSearchResponse, error) {
	args := m.Called(c, hospitalID, req)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*dto.PatientSearchResponse), args.Error(1)
}
