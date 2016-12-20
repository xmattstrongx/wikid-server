package services

// AccountService manages accounts.
type AccountService interface{}

type accountService struct{}

// NewAccountService constructs a new account service.
var NewAccountService = func() AccountService {
	return &accountService{}
}
