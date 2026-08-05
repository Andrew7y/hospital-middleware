package main

import (
	"hospital-middleware/internal/config"
	"hospital-middleware/internal/handler"
	"hospital-middleware/internal/repository"
	"hospital-middleware/internal/router"
	"hospital-middleware/internal/service"
)

func main() {
	configuration := config.LoadConfig()
	db := config.ConnectDatabase(configuration)
	config.MigrateDatabase(db)

	staffRepository := repository.NewStaffRepository(db)
	staffService := service.NewStaffService(staffRepository, configuration.JwtSecret)
	staffHandler := handler.NewStaffHandler(staffService)

	patientRepository := repository.NewPatientRepository(db)
	patientService := service.NewPatientService(patientRepository)
	patientHandler := handler.NewPatientHandler(patientService)

	r := router.SetupRouter(
		staffHandler,
		patientHandler,
		configuration.JwtSecret,
	)

	err := r.Run(`:8080`)
	if err != nil {
		return
	}
}
