package tests

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"hospital-middleware/internal/dto"
	"hospital-middleware/internal/handler"
	"hospital-middleware/internal/middleware"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchPatient_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockPatientService := new(MockPatientService)
	patientHandler := handler.NewPatientHandler(mockPatientService)

	router := gin.New()

	router.GET("/patient/search", func(c *gin.Context) {
		c.Set("hospital_id", uint(1))
		patientHandler.SearchPatient(c)
	})

	reqBody := dto.PatientSearchRequest{
		Page:      1,
		PageSize:  20,
		SortBy:    "first_name_en",
		SortOrder: "asc",
	}

	nationalID := "1234567890123"
	passportId := ""
	expectedResponse := &dto.PatientSearchResponse{
		Patients: []dto.PatientSearchResult{
			{
				ID:           1,
				HospitalID:   1,
				FirstNameTH:  "สมชาย",
				MiddleNameTH: "",
				LastNameTH:   "ใจดี",
				FirstNameEN:  "John",
				MiddleNameEN: "",
				LastNameEN:   "Doe",
				DateOfBirth:  "01/01/1990",
				PatientHN:    "HN001",
				NationalID:   &nationalID,
				PassportID:   &passportId,
				PhoneNumber:  "0812345678",
				Email:        "john@example.com",
				Gender:       "Male",
			},
		},
		Pagination: dto.PaginationMeta{
			Page:       1,
			PageSize:   20,
			TotalItems: 1,
			TotalPages: 1,
		},
		Sorting: dto.SortingMeta{
			SortBy:    "first_name_en",
			SortOrder: "asc",
		},
	}

	mockPatientService.On(
		"SearchPatient",
		mock.Anything,
		uint(1),
		reqBody,
	).Return(
		expectedResponse,
		nil,
	).Once()

	req := httptest.NewRequest(
		http.MethodGet,
		"/patient/search?page=1&page_size=20&sort_by=first_name_en&sort_order=asc",
		nil,
	)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(
		t,
		http.StatusOK,
		recorder.Code,
	)

	var response struct {
		Message string                    `json:"message"`
		Data    dto.PatientSearchResponse `json:"data"`
	}

	err := json.Unmarshal(
		recorder.Body.Bytes(),
		&response,
	)

	require.NoError(t, err)

	assert.Equal(
		t,
		"Patients retrieved successfully",
		response.Message,
	)

	require.Len(
		t,
		response.Data.Patients,
		1,
	)

	assert.Equal(
		t,
		uint(1),
		response.Data.Patients[0].ID,
	)

	assert.Equal(
		t,
		"John",
		response.Data.Patients[0].FirstNameEN,
	)

	assert.Equal(
		t,
		"Doe",
		response.Data.Patients[0].LastNameEN,
	)

	assert.Equal(
		t,
		1,
		response.Data.Pagination.Page,
	)

	assert.Equal(
		t,
		20,
		response.Data.Pagination.PageSize,
	)

	assert.Equal(
		t,
		int64(1),
		response.Data.Pagination.TotalItems,
	)

	assert.Equal(
		t,
		1,
		response.Data.Pagination.TotalPages,
	)

	assert.Equal(
		t,
		"first_name_en",
		response.Data.Sorting.SortBy,
	)

	assert.Equal(
		t,
		"asc",
		response.Data.Sorting.SortOrder,
	)

	mockPatientService.AssertExpectations(t)
}

func TestSearchPatient_Failed_RequireLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockPatientService := new(MockPatientService)
	patientHandler := handler.NewPatientHandler(mockPatientService)

	router := gin.New()

	router.GET(
		"/patient/search",
		middleware.JWTAuth("test-secret"),
		patientHandler.SearchPatient,
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/patient/search?page=1&page_size=20",
		nil,
	)

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

	err := json.Unmarshal(
		recorder.Body.Bytes(),
		&response,
	)

	require.NoError(t, err)

	assert.Equal(
		t,
		"Authorization header is missing",
		response.Error,
	)

	mockPatientService.AssertNotCalled(
		t,
		"SearchPatient",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	)
}
