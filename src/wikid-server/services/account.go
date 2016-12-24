package services

import (
	"crypto/rand"
	"errors"
	"net/mail"
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

// -------------------------------------------------------------------------- //

func NewAccountService() IAccountService {
	return &accountService{
		accountRepository: repositories.NewAccountRepository(),
	}
}

func (this *accountService) CreateAccount(email, password string) (*data.Account, error) {
	// Validate email and password.
	if err := validateEmail(email); err != nil {
		return nil, err
	}

	if err := validatePassword(password); err != nil {
		return nil, err
	}

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
		ID:          uuid.NewV4().String(),
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

func validateEmail(email string) error {
	address, err := mail.ParseAddress(email)
	if err != nil || address.Name != "" {
		return errors.New("Invalid email address.")
	}

	return nil
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("Password must be at last 8 characters in length.")
	}

	return nil
}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 64)
	_, err := rand.Read(salt)
	return salt, err
}

func hashPassword(password, salt []byte) ([]byte, error) {
	saltedPassword := append(password, salt...)
	return bcrypt.GenerateFromPassword(saltedPassword, 10)
}
