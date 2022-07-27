package entities

import (
	"errors"
	"testing"
)

func TestAccountValidateSuccess(t *testing.T) {
	account := Account{
		Name:    "test",
		CPF:     "89414768017",
		Secret:  "u69UFv$*ETH4",
		Balance: 100,
	}
	result := account.Validate()

	if result != nil {
		t.Errorf("result %v, expected %v", result, nil)
	}
}

func TestAccountValidateEmptyName(t *testing.T) {
	account := Account{
		CPF:     "89414768017",
		Secret:  "u69UFv$*ETH4",
		Balance: 100,
	}
	result := account.Validate()
	expected := errors.New("name can't be empty")

	if result.Error() != expected.Error() {
		t.Errorf("result %v, expected %v", result, expected)
	}
}

func TestAccountValidateEmptyCPF(t *testing.T) {
	account := Account{
		Name:    "test",
		Secret:  "u69UFv$*ETH4",
		Balance: 100,
	}
	result := account.Validate()
	expected := errors.New("CPF can't be empty")

	if result.Error() != expected.Error() {
		t.Errorf("result %v, expected %v", result, expected)
	}
}

func TestAccountValidateLenCPF(t *testing.T) {
	account := Account{
		Name:    "test",
		CPF:     "8941476801",
		Secret:  "u69UFv$*ETH4",
		Balance: 100,
	}
	result := account.Validate()
	expected := errors.New("CPF must be 11 characters long")

	if result.Error() != expected.Error() {
		t.Errorf("result %v, expected %v", result, expected)
	}
}

func TestAccountValidateEmptySecret(t *testing.T) {
	account := Account{
		Name:    "test",
		CPF:     "89414768017",
		Balance: 100,
	}
	result := account.Validate()
	expected := errors.New("secret can't be empty")

	if result.Error() != expected.Error() {
		t.Errorf("result %v, expected %v", result, expected)
	}
}

func TestAccountValidateLenSecret(t *testing.T) {
	account := Account{
		Name:    "test",
		CPF:     "89414768017",
		Secret:  "*ETH4",
		Balance: 100,
	}
	result := account.Validate()
	expected := errors.New("secret must be at least 8 characters long")

	if result.Error() != expected.Error() {
		t.Errorf("result %v, expected %v", result, expected)
	}
}

func TestAccountGenerateHashSuccess(t *testing.T) {
	account := Account{
		Name:    "test",
		CPF:     "89414768017",
		Secret:  "u69UFv$*ETH4",
		Balance: 100,
	}
	result := account.GenerateHash()

	if result != nil {
		t.Errorf("result %v, expected %v", result, nil)
	}
}

func TestAccountIsCorrectSecretSuccess(t *testing.T) {
	accountDb := Account{
		Name:    "test",
		CPF:     "89414768017",
		Secret:  "$2a$10$kf6gYk/3/84cg9yvNEGkrOVWchqFOcKjJ2a/dax.BZbdDhklbX6gS",
		Balance: 100,
	}
	secretResult := "u69UFv$*ETH4"

	isCorrect := accountDb.IsCorrectSecret(secretResult)

	if !isCorrect {
		t.Errorf("result %t, expected %t", false, true)
	}
}

func TestAccountIsCorrectSecretFailure(t *testing.T) {
	accountDb := Account{
		Name:    "test",
		CPF:     "89414768017",
		Secret:  "$2a$10$kf6gYk/3/84cg9yvNEGkrOVWchqFOcKjJ2a/dax.BZbdDhklbX6gS",
		Balance: 100,
	}
	secretResult := "u69UFv$*ETH"

	isCorrect := accountDb.IsCorrectSecret(secretResult)

	if isCorrect {
		t.Errorf("result %t, expected %t", false, true)
	}
}
