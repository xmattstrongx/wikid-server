package repositories

import "time"

type AccountEntity struct {
	ID          string    `db:"id"`
	Email       string    `db:"email"`
	Password    []byte    `db:"password"`
	Salt        []byte    `db:"salt"`
	CreatedTime time.Time `db:"created_time"`
	UpdatedTime time.Time `db:"updated_time"`
	DeletedTime time.Time `db:"deleted_time"`
}

type accountRepository struct{}

type IAccountRepository interface {
	CreateAccount(account *AccountEntity) error
}

func NewAccountRepository() IAccountRepository {
	return &accountRepository{}
}

// -------------------------------------------------------------------------- //

func (this *accountRepository) CreateAccount(account *AccountEntity) error {
	_, err := _db.NamedExec(`
		INSERT INTO account ( id,  email,  password,  salt,  created_time)
		values              (:id, :email, :password, :salt, :created_time);`,
		account)

	return err
}
