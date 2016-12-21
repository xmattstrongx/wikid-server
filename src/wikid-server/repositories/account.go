package repositories

import "wikid-server/models"

type IAccountRepository interface {
	CreateAccount(account *models.Account) error
}

type accountRepository struct{}

// -------------------------------------------------------------------------- //

func NewAccountRepository() IAccountRepository {
	return &accountRepository{}
}

func (this *accountRepository) CreateAccount(account *models.Account) error {
	result, err := _db.NamedExec(`
		INSERT INTO account ( email,  password,  salt,  created_time)
		values              (:email, :password, :salt, :created_time);`,
		account)

	if err != nil {
		return err
	}

	account.Id, err = result.LastInsertId()

	return err
}
