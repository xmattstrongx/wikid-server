package services

import (
	"wikid-server/app"

	"github.com/jmoiron/sqlx"
)

type IHealthService interface {
	CheckDatabase() error
}

type healthService struct {
	db *sqlx.DB
}

// -------------------------------------------------------------------------- //

func NewHealthService() IHealthService {
	return &healthService{
		db: app.GetDatabase(),
	}
}

func (this *healthService) CheckDatabase() error {
	return this.db.Ping()
}
