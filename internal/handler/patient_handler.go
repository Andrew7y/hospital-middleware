package handler

import (
	"github.com/gin-gonic/gin"
	"hospital-middleware/internal/dto"
	"hospital-middleware/internal/service"
	"net/http"
)

type PatientHandler struct {
	patientService service.PatientService
}

func NewPatientHandler(patientService service.PatientService) *PatientHandler {
	return &PatientHandler{patientService: patientService}
}

func (h *PatientHandler) SearchPatient(c *gin.Context) {
	var req dto.PatientSearchRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hospitalID, exists := c.Get("hospital_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized hospital_id not found",
		})
		return
	}

	res, err := h.patientService.SearchPatient(
		c.Request.Context(),
		hospitalID.(uint),
		req,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Patients retrieved successfully",
		"data":    res,
	})
}
