package services

import (
	"wikid-server/models"
	"wikid-server/repositories"
)

type ISessionService interface {
	CreateSession(accountId string) (*models.Session, error)
}

type sessionService struct {
	sessionRepository repositories.ISessionRepository
}

// -------------------------------------------------------------------------- //

func NewSessionService() ISessionService {
	return &sessionService{
		sessionRepository: repositories.SessionRepository,
	}
}

func (this *sessionService) CreateSession(accountId string) (*models.Session, error) {
	return this.sessionRepository.CreateSession(accountId)
}
