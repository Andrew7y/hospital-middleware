package dto

type StaffCreateRequest struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	HospitalID uint   `json:"hospital_id" binding:"required"`
}

type StaffLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type StaffResponse struct {
	ID         uint   `json:"id"`
	Username   string `json:"username"`
	HospitalID uint   `json:"hospital_id"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   string `json:"expires_at"`
}
