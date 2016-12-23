package repositories

import (
	"wikid-server/app"
	"wikid-server/models/data"

	"github.com/jmoiron/sqlx"
)

type IAccountRepository interface {
	CreateAccount(account *data.Account) error
}

type accountRepository struct {
	db *sqlx.DB
}

// -------------------------------------------------------------------------- //

func NewAccountRepository() IAccountRepository {
	return &accountRepository{
		db: app.GetDatabase(),
	}
}

func (this *accountRepository) CreateAccount(account *data.Account) error {
	_, err := this.db.NamedExec(`
		INSERT INTO account ( id,  email,  password,  salt,  created_time)
		values              (:id, :email, :password, :salt, :created_time);`,
		account)

	return err
}
