package services

import (
	"crypto/rand"
	"time"
	"wikid-server/models/data"
	"wikid-server/repositories"

	uuid "github.com/satori/go.uuid"

	"golang.org/x/crypto/bcrypt"
)

type IAccountService interface {
	CreateAccount(email, password string) (*data.Account, error)
}

type accountService struct {
	accountRepository repositories.IAccountRepository
}

var NewAccountService = func() IAccountService {
	return &accountService{
		accountRepository: repositories.NewAccountRepository(),
	}
}

// -------------------------------------------------------------------------- //

func (this *accountService) CreateAccount(email, password string) (*data.Account, error) {
	// TODO: Validate account values.

	// Generate salt and hash password.
	salt, err := generateSalt()
	if err != nil {
		return nil, err
	}

	hashedPassword, err := hashPassword([]byte(password), salt)
	if err != nil {
		return nil, err
	}

	// Create account.
	account := &data.Account{
		Id:          uuid.NewV4().String(),
		Email:       email,
		Password:    hashedPassword,
		Salt:        salt,
		CreatedTime: time.Now(),
	}

	if err := this.accountRepository.CreateAccount(account); err != nil {
		return nil, err
	}

	return account, nil
}

// -------------------------------------------------------------------------- //

func generateSalt() ([]byte, error) {
	salt := make([]byte, 64)
	_, err := rand.Read(salt)
	return salt, err
}

func hashPassword(password, salt []byte) ([]byte, error) {
	saltedPassword := append(password, salt...)
	return bcrypt.GenerateFromPassword(saltedPassword, 10)
}
