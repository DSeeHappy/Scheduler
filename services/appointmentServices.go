package services

import "Scheduler/repositories"

type AppointmentService struct {
	repository *repositories.AppointmentRepository
}

func NewAppointmentService(repository *repositories.AppointmentRepository) *AppointmentService {
	return &AppointmentService{
		repository: repository,
	}
}
