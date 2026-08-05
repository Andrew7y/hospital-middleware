package token

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateToken(staffID, hospitalID uint, username, jwtSecret string) (string, error) {
	claims := &StaffClaims{
		StaffID:    staffID,
		HospitalID: hospitalID,
		Username:   username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 3)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}
