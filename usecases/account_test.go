package usecases

import (
	"github.com/lucianetedesco/banking-api/entities"
	"reflect"
	"testing"
	"time"
)

func TestCreateAccount(t *testing.T) {
	account := entities.Account{
		Name:    "test",
		CPF:     "89414768017",
		Secret:  "u69UFv$*ETH4",
		Balance: 100,
	}

	expected := uint(1)
	useCaseAccount := NewAccountUseCase(AccountRepositoryMock{})
	result, err := useCaseAccount.SaveAccount(account)

	if result != expected && err != nil {
		t.Errorf("result %v, expected %v", result, expected)
	}
}

func TestGetAllAccount(t *testing.T) {
	expected := []entities.Account{
		{
			ID:        1,
			Name:      "test-1",
			CPF:       "89414768017",
			Secret:    "u69UFv$*ETH4",
			Balance:   100,
			CreatedAt: time.Time{},
		},
		{
			ID:        2,
			Name:      "test-2",
			CPF:       "32894658028",
			Secret:    "pQ*Lr8&7",
			Balance:   1500,
			CreatedAt: time.Time{},
		},
	}

	useCaseAccount := NewAccountUseCase(AccountRepositoryMock{})
	result, err := useCaseAccount.GetAllAccounts()

	if !reflect.DeepEqual(result, expected) && err != nil {
		t.Errorf("result %v, expected %v", result, expected)
	}
}

func TestGetBalanceAccount(t *testing.T) {
	expected := 100.0

	accountId := uint(1)
	useCaseAccount := NewAccountUseCase(AccountRepositoryMock{})
	result, err := useCaseAccount.GetBalanceAccount(accountId)

	if result != expected && err != nil {
		t.Errorf("result %v, expected %v", result, expected)
	}
}

func TestGetAccount(t *testing.T) {
	expected :=
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoxLCJleHAiOjE2NTg5NzYxNTJ9.uMihMTsHVWBrkT5rnZg7BISyIx1sULt1f4LjqylGz8k"

	login := entities.Login{
		CPF:    "89414768017",
		Secret: "u69UFv$*ETH4",
	}

	useCaseAccount := NewAccountUseCase(AccountRepositoryMock{})
	result, err := useCaseAccount.GetAccount(login)

	if !reflect.DeepEqual(result, expected) && err != nil {
		t.Errorf("result %v, expected %v", result, expected)
	}
}
