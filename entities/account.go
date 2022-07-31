package entities

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Account struct {
	ID        uint      `json:"id,omitempty"`
	Name      string    `json:"name" gorm:"type:varchar(100);not null"`
	CPF       string    `json:"cpf" gorm:"type:varchar(11);unique;not null"`
	Secret    string    `json:"secret" gorm:"type:varchar(60);not null"`
	Balance   float64   `json:"balance" gorm:"default:0;not null"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type Login struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

func (a Account) Validate() error {
	if a.Name == "" {
		return errors.New("name can't be empty")
	}

	if a.CPF == "" {
		return errors.New("CPF can't be empty")
	}

	if len(a.CPF) != 11 {
		return errors.New("CPF must be 11 characters long")
	}

	if a.Secret == "" {
		return errors.New("secret can't be empty")
	}

	if len(a.Secret) < 8 {
		return errors.New("secret must be at least 8 characters long")
	}

	return nil
}

func (a *Account) GenerateHash() error {
	secret, err := bcrypt.GenerateFromPassword([]byte(a.Secret), bcrypt.DefaultCost)

	if err != nil {
		return errors.New("could not create account hash")
	}

	a.Secret = string(secret)

	return nil
}

func (a *Account) IsCorrectSecret(secret string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(a.Secret), []byte(secret))
	return err == nil
}
