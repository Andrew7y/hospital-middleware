package router

import (
	"github.com/gin-gonic/gin"
	"hospital-middleware/internal/handler"
	"hospital-middleware/internal/middleware"
)

func SetupRouter(
	staffHandler *handler.StaffHandler,
	patientHandler *handler.PatientHandler,
	jwtSecret string,
) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/staff/create", staffHandler.CreateStaff)
		api.POST("/staff/login", staffHandler.Login)

		protected := api.Group("")
		protected.Use(middleware.JWTAuth(jwtSecret))
		{
			protected.GET("/patient/search", patientHandler.SearchPatient)
		}
	}

	return r
}
