package usecases

import (
	"github.com/lucianetedesco/banking-api/entities"
	"time"
)

type AccountRepositoryMock struct {
}

func (m AccountRepositoryMock) SaveAccount(account *entities.Account) (uint, error) {
	return 1, nil
}

func (m AccountRepositoryMock) GetAllAccounts() ([]entities.Account, error) {
	accounts := []entities.Account{
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
	return accounts, nil
}
func (m AccountRepositoryMock) GetAccountByCPF(cpf string) (entities.Account, error) {
	account := entities.Account{
		ID:        1,
		Name:      "test-1",
		CPF:       "89414768017",
		Secret:    "$2a$10$LpB41uXk9am/FlXUy1JfeeIcl/skhTorp7Mi51TJQ3ouel5oV9UFq",
		Balance:   100,
		CreatedAt: time.Time{},
	}
	return account, nil
}

func (m AccountRepositoryMock) GetBalanceAccount(accountId uint) (float64, error) {
	return 100.0, nil
}

func (m AccountRepositoryMock) UpdateBalanceAccount(accountID uint, newBaLance float64) error {
	return nil
}
