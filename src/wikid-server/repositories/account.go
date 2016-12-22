package repositories

import "wikid-server/models"

type IAccountRepository interface {
	CreateAccount(account *models.Account) error
}

type accountRepository struct{}

var NewAccountRepository = func() IAccountRepository {
	return &accountRepository{}
}

// -------------------------------------------------------------------------- //

func (this *accountRepository) CreateAccount(account *models.Account) error {
	_, err := _db.NamedExec(`
		INSERT INTO account ( id,  email,  password,  salt,  created_time)
		values              (:id, :email, :password, :salt, :created_time);`,
		account)

	return err
}
