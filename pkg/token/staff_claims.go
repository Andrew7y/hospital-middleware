package token

import "github.com/golang-jwt/jwt/v5"

type StaffClaims struct {
	StaffID    uint   `json:"staff_id"`
	HospitalID uint   `json:"hospital_id"`
	Username   string `json:"username"`
	jwt.RegisteredClaims
}
