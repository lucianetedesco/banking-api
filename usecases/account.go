package usecases

import (
	"errors"
	"github.com/lucianetedesco/banking-api/core"
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

func (u AccountUseCase) GetBalanceAccount(accountId uint) (float64, error) {
	return u.AccountRepository.GetBalanceAccount(accountId)
}

func (u AccountUseCase) GetAccount(login entities.Login) (string, error) {
	account, err := u.AccountRepository.GetAccountByCPF(login.CPF)
	if err != nil {
		return "", err
	}

	if isCorrectSecret := account.IsCorrectSecret(login.Secret); !isCorrectSecret {
		return "", errors.New("incorrect secret")
	}

	token, err := core.GenerateToken(account.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
