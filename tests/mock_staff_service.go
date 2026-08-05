package tests

import (
	"context"
	"github.com/stretchr/testify/mock"
	"hospital-middleware/internal/dto"
)

type MockStaffService struct {
	mock.Mock
}

func (m *MockStaffService) CreateStaff(
	c context.Context,
	req dto.StaffCreateRequest,
) (*dto.StaffResponse, error) {
	args := m.Called(c, req)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*dto.StaffResponse), args.Error(1)
}

func (m *MockStaffService) Login(
	c context.Context,
	req dto.StaffLoginRequest,
) (*dto.LoginResponse, error) {
	args := m.Called(c, req)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*dto.LoginResponse), args.Error(1)
}
