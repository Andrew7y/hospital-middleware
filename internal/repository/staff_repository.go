package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"hospital-middleware/internal/model"
)

type StaffRepository interface {
	CreateStaff(c context.Context, staff *model.Staff) error
	FindByUsername(c context.Context, username string) (*model.Staff, error)
}

type staffRepository struct {
	db *gorm.DB
}

func NewStaffRepository(db *gorm.DB) StaffRepository {
	return &staffRepository{db: db}
}

func (r *staffRepository) CreateStaff(c context.Context, staff *model.Staff) error {
	return r.db.WithContext(c).Create(staff).Error
}

func (r *staffRepository) FindByUsername(c context.Context, username string) (*model.Staff, error) {
	var staff model.Staff
	err := r.db.WithContext(c).Where("username = ?", username).First(&staff).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &staff, nil
}
