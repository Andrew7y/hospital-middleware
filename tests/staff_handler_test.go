package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"hospital-middleware/internal/dto"
	"hospital-middleware/internal/handler"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateStaff_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockStaffService := new(MockStaffService)
	router := gin.New()
	staffHaandler := handler.NewStaffHandler(mockStaffService)

	router.POST("/api/staff/create", staffHaandler.CreateStaff)

	reqbody := dto.StaffCreateRequest{
		Username:   "testuser",
		Password:   "12345",
		HospitalID: 1,
	}

	expectedResponse := dto.StaffResponse{
		ID:         1,
		Username:   "testuser",
		HospitalID: 1,
	}

	mockStaffService.On(
		"CreateStaff",
		mock.Anything,
		reqbody,
	).Return(
		&expectedResponse,
		nil,
	).Once()

	body, err := json.Marshal(reqbody)
	require.NoError(t, err)

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/staff/create",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)

	var response struct {
		Message string            `json:"message"`
		Data    dto.StaffResponse `json:"data"`
	}

	err = json.Unmarshal(
		recorder.Body.Bytes(),
		&response,
	)

	require.NoError(t, err)

	assert.Equal(
		t,
		"Staff created successfully. Please log in.",
		response.Message,
	)

	assert.Equal(
		t,
		uint(1),
		response.Data.ID,
	)

	assert.Equal(
		t,
		"testuser",
		response.Data.Username,
	)

	assert.Equal(
		t,
		uint(1),
		response.Data.HospitalID,
	)

	mockStaffService.AssertExpectations(t)
}

func TestCreateStaff_Failed_UsernameAlreadyExist(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockStaffService := new(MockStaffService)
	staffHandler := handler.NewStaffHandler(mockStaffService)

	router := gin.New()
	router.POST("/staff/create", staffHandler.CreateStaff)

	reqBody := dto.StaffCreateRequest{
		Username:   "testuser",
		Password:   "12345",
		HospitalID: 1,
	}

	mockStaffService.On(
		"CreateStaff",
		mock.Anything,
		reqBody,
	).Return(
		nil,
		errors.New("username already exists"),
	).Once()

	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest(
		http.MethodPost,
		"/staff/create",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(
		t,
		http.StatusConflict,
		recorder.Code,
	)

	var response struct {
		Error string `json:"error"`
	}

	err = json.Unmarshal(
		recorder.Body.Bytes(),
		&response,
	)

	require.NoError(t, err)

	assert.Equal(
		t,
		"username already exists",
		response.Error,
	)

	mockStaffService.AssertExpectations(t)
}

func TestCreateStaff_Failed_EmptyUsernameAndPassword(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockStaffService := new(MockStaffService)
	staffHandler := handler.NewStaffHandler(mockStaffService)

	router := gin.New()
	router.POST("/staff/create", staffHandler.CreateStaff)

	reqBody := dto.StaffCreateRequest{
		Username:   "",
		Password:   "",
		HospitalID: 1,
	}

	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest(
		http.MethodPost,
		"/staff/create",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(
		t,
		http.StatusBadRequest,
		recorder.Code,
	)

	mockStaffService.AssertNotCalled(
		t,
		"CreateStaff",
		mock.Anything,
		mock.Anything,
	)
}

func TestLogin_Success_GotJWTToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockStaffService := new(MockStaffService)
	staffHandler := handler.NewStaffHandler(mockStaffService)

	router := gin.New()
	router.POST("/staff/login", staffHandler.Login)

	reqBody := dto.StaffLoginRequest{
		Username: "testuser",
		Password: "123456",
	}

	expectedResponse := &dto.LoginResponse{
		AccessToken: "mock-jwt-token",
		ExpiresAt:   "2026-08-09T12:00:00Z",
	}

	mockStaffService.On(
		"Login",
		mock.Anything,
		reqBody,
	).Return(
		expectedResponse,
		nil,
	).Once()

	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest(
		http.MethodPost,
		"/staff/login",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(
		t,
		http.StatusOK,
		recorder.Code,
	)

	var response struct {
		Message string            `json:"message"`
		Data    dto.LoginResponse `json:"data"`
	}
	err = json.Unmarshal(
		recorder.Body.Bytes(),
		&response,
	)

	require.NoError(t, err)

	assert.Equal(
		t,
		"Login successful",
		response.Message,
	)

	assert.Equal(
		t,
		"mock-jwt-token",
		response.Data.AccessToken,
	)

	assert.NotEmpty(
		t,
		response.Data.AccessToken,
	)

	assert.Equal(
		t,
		"2026-08-09T12:00:00Z",
		response.Data.ExpiresAt,
	)

	mockStaffService.AssertExpectations(t)
}

func TestLogin_Failed_IncorrectPassword(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockStaffService := new(MockStaffService)
	staffHandler := handler.NewStaffHandler(mockStaffService)

	router := gin.New()
	router.POST("/staff/login", staffHandler.Login)

	reqBody := dto.StaffLoginRequest{
		Username: "testuser",
		Password: "wrong-password",
	}

	mockStaffService.
		On(
			"Login",
			mock.Anything,
			reqBody,
		).Return(
		nil,
		errors.New("incorrect password"),
	).Once()

	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest(
		http.MethodPost,
		"/staff/login",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(
		t,
		http.StatusUnauthorized,
		recorder.Code,
	)

	var response struct {
		Error string `json:"error"`
	}

	err = json.Unmarshal(
		recorder.Body.Bytes(),
		&response,
	)

	require.NoError(t, err)

	assert.Equal(
		t,
		"incorrect password",
		response.Error,
	)

	mockStaffService.AssertExpectations(t)
}

func TestLogin_Failed_UsernameNotExist(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockStaffService := new(MockStaffService)
	staffHandler := handler.NewStaffHandler(mockStaffService)

	router := gin.New()
	router.POST("/staff/login", staffHandler.Login)

	reqBody := dto.StaffLoginRequest{
		Username: "unknown-user",
		Password: "123456",
	}

	mockStaffService.On(
		"Login",
		mock.Anything,
		reqBody,
	).Return(
		nil,
		errors.New("username does not exist"),
	).Once()

	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest(
		http.MethodPost,
		"/staff/login",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(
		t,
		http.StatusUnauthorized,
		recorder.Code,
	)

	var response struct {
		Error string `json:"error"`
	}

	err = json.Unmarshal(
		recorder.Body.Bytes(),
		&response,
	)

	require.NoError(t, err)

	assert.Equal(
		t,
		"username does not exist",
		response.Error,
	)

	mockStaffService.AssertExpectations(t)
}
