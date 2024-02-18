package controllers

import "Scheduler/services"

type AppointmentController struct {
	service *services.AppointmentService
}

func NewAppointmentsController(service *services.AppointmentService) *AppointmentController {
	return &AppointmentController{
		service: service,
	}
}
