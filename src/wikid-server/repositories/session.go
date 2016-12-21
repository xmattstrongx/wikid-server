package repositories

import "wikid-server/models"

type ISessionRepository interface {
	CreateSession(accountId string) (*models.Session, error)
}

type sessionRepository struct{}

// -------------------------------------------------------------------------- //

func NewSessionRepository() ISessionRepository {
	return &sessionRepository{}
}

func (this *sessionRepository) CreateSession(accountId string) (*models.Session, error) {
	return nil, nil
}
