package handler

import (
	"github.com/gin-gonic/gin"
	"hospital-middleware/internal/dto"
	"hospital-middleware/internal/service"
	"net/http"
)

type StaffHandler struct {
	staffService service.StaffService
}

func NewStaffHandler(staffService service.StaffService) *StaffHandler {
	return &StaffHandler{staffService: staffService}
}

func (h *StaffHandler) CreateStaff(c *gin.Context) {
	var req dto.StaffCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.staffService.CreateStaff(c.Request.Context(), req)
	if err != nil {
		if err.Error() == "username already exists" {
			c.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create staff",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Staff created successfully. Please log in.",
		"data":    res,
	})
}

func (h *StaffHandler) Login(c *gin.Context) {
	var req dto.StaffLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.staffService.Login(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"data":    res,
	})
}
