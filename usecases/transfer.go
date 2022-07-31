package usecases

import (
	"errors"
	"github.com/lucianetedesco/banking-api/core"
	"github.com/lucianetedesco/banking-api/entities"
	"github.com/lucianetedesco/banking-api/repositories"
	"strconv"
)

type TransferUseCase struct {
	TransferRepository repositories.TransferRepository
}

func NewTransferUseCase(r repositories.TransferRepository) *TransferUseCase {
	return &TransferUseCase{TransferRepository: r}
}

func (u TransferUseCase) SaveTransfer(transfer entities.Transfer, token string) (uint, error) {
	accountOriginID, err := core.GetAccountID(token)
	if err != nil {
		return 0, errors.New("user unauthorized")
	}

	u64, _ := strconv.ParseUint(accountOriginID, 10, 32)
	wd := uint(u64)

	transfer.AccountOriginId = wd

	d := core.GetDatabaseConnectionInstance()
	tx := d.Db.Begin()

	repositoryAccount := repositories.NewAccountRepository(tx)
	useCaseAccount := NewAccountUseCase(repositoryAccount)

	balanceAccountOrigin, err := useCaseAccount.GetBalanceAccount(transfer.AccountOriginId)
	if err != nil {
		return 0, err
	}
	if balanceAccountOrigin < transfer.Amount {
		return 0, errors.New("insufficient balance for transfer")
	}

	balanceAccountDestination, err := useCaseAccount.GetBalanceAccount(transfer.AccountDestinationId)
	if err != nil {
		return 0, err
	}

	err = repositoryAccount.UpdateBalanceAccount(transfer.AccountOriginId, balanceAccountOrigin-transfer.Amount)
	err = repositoryAccount.UpdateBalanceAccount(transfer.AccountDestinationId, balanceAccountDestination+transfer.Amount)
	id, err := u.TransferRepository.SaveTransfer(&transfer)

	tx.Commit()
	return id, err
}

func (u TransferUseCase) GetTransfers(token string) ([]entities.Transfer, error) {
	accountOriginID, err := core.GetAccountID(token)
	if err != nil {
		return nil, errors.New("user unauthorized")
	}

	return u.TransferRepository.GetAllTransfers(accountOriginID)
}
