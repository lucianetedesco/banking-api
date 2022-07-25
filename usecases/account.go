package usecases

import (
	"github.com/lucianetedesco/banking-api/entities"
	"github.com/lucianetedesco/banking-api/repositories"
)

type AccountUseCase struct {
	AccountRepository repositories.AccountRepository
}

func NewAccountUseCase(r repositories.AccountRepository) *AccountUseCase {
	return &AccountUseCase{AccountRepository: r}
}

func (u AccountUseCase) SaveAccount(account entities.Account) (uint, error) {
	if err := account.GenerateHash(); err != nil {
		return 0, err
	}

	id, err := u.AccountRepository.SaveAccount(&account)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u AccountUseCase) GetAllAccounts() ([]entities.Account, error) {
	return u.AccountRepository.GetAllAccounts()
}
