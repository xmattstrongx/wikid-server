package services

import "wikid-server/repositories"

type healthService struct{}

type IHealthService interface {
	CheckDatabase() error
}

func NewHealthService() IHealthService {
	return &healthService{}
}

// -------------------------------------------------------------------------- //

func (this *healthService) CheckDatabase() error {
	return repositories.PingDatabase()
}
