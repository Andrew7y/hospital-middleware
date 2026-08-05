package service

import (
	"context"
	"errors"
	"hospital-middleware/internal/dto"
	"hospital-middleware/internal/model"
	"hospital-middleware/internal/repository"
	"hospital-middleware/pkg/password"
	"hospital-middleware/pkg/token"
	"time"
)

type StaffService interface {
	CreateStaff(c context.Context, req dto.StaffCreateRequest) (*dto.StaffResponse, error)
	Login(c context.Context, req dto.StaffLoginRequest) (*dto.LoginResponse, error)
}

type staffService struct {
	staffRepo repository.StaffRepository
	jwtSecret string
}

func NewStaffService(repo repository.StaffRepository, jwtSecret string) StaffService {
	return &staffService{
		staffRepo: repo,
		jwtSecret: jwtSecret,
	}
}

func (s *staffService) CreateStaff(c context.Context, req dto.StaffCreateRequest) (*dto.StaffResponse, error) {
	existing, err := s.staffRepo.FindByUsername(c, req.Username)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, errors.New("username already exists")
	}

	hashedPassword, err := password.HashedPassword(req.Password)
	if err != nil {
		return nil, err
	}

	newStaff := &model.Staff{
		Username:   req.Username,
		Password:   hashedPassword,
		HospitalID: req.HospitalID,
	}

	if err := s.staffRepo.CreateStaff(c, newStaff); err != nil {
		return nil, err
	}

	return &dto.StaffResponse{
		ID:         newStaff.ID,
		Username:   newStaff.Username,
		HospitalID: newStaff.HospitalID,
	}, nil
}

func (s *staffService) Login(c context.Context, req dto.StaffLoginRequest) (*dto.LoginResponse, error) {
	staff, err := s.staffRepo.FindByUsername(c, req.Username)
	if err != nil {
		return nil, err
	}
	if staff == nil {
		return nil, errors.New("username does not exist")
	}

	if err := password.ComparePassword(staff.Password, req.Password); err != nil {
		return nil, errors.New("incorrect password")
	}

	expriationTime := time.Now().Add(24 * time.Hour)
	tokenString, err := token.GenerateToken(
		staff.ID,
		staff.HospitalID,
		staff.Username,
		s.jwtSecret,
		expriationTime,
	)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken: tokenString,
		ExpiresAt:   expriationTime.Format(time.RFC3339),
	}, nil

}
