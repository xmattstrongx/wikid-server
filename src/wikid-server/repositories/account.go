package repositories

import "wikid-server/models/data"

type IAccountRepository interface {
	CreateAccount(account *data.Account) error
}

type accountRepository struct{}

var NewAccountRepository = func() IAccountRepository {
	return &accountRepository{}
}

// -------------------------------------------------------------------------- //

func (this *accountRepository) CreateAccount(account *data.Account) error {
	_, err := _db.NamedExec(`
		INSERT INTO account ( id,  email,  password,  salt,  created_time)
		values              (:id, :email, :password, :salt, :created_time);`,
		account)

	return err
}
