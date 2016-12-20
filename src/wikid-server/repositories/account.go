package repositories

// AccountRepository is a repository for accounts.
type AccountRepository interface{}

type accountRepository struct{}

// NewAccountRepository constructs a new account repository.
var NewAccountRepository = func() AccountRepository {
	return &accountRepository{}
}
