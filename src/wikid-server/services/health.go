package services

import "wikid-server/repositories"

type healthService struct{}

type IHealthService interface {
	GenerateReport() (*HealthReport, error)
}

func NewHealthService() IHealthService {
	return &healthService{}
}

// -------------------------------------------------------------------------- //

type HealthReport struct {
	DatabaseHealthy bool
}

// GenerateReport generates a health report of the server's external
// dependencies.
func (this *healthService) GenerateReport() (*HealthReport, error) {
	healthReport := &HealthReport{
		DatabaseHealthy: true,
	}

	if err := repositories.PingDatabase(); err != nil {
		healthReport.DatabaseHealthy = false
	}

	return healthReport, nil
}
