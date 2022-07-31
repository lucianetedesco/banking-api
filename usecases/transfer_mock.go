package usecases

import (
	"github.com/lucianetedesco/banking-api/entities"
	"time"
)

type TransferRepositoryMock struct {
}

func (m TransferRepositoryMock) SaveTransfer(transfer *entities.Transfer) (uint, error) {
	return 1, nil
}
func (m TransferRepositoryMock) GetAllTransfers(token string) ([]entities.Transfer, error) {
	transfers := []entities.Transfer{
		{
			ID:                   1,
			AccountOriginId:      1,
			AccountDestinationId: 2,
			Amount:               100,
			CreatedAt:            time.Time{},
		},
		{
			ID:                   2,
			AccountOriginId:      1,
			AccountDestinationId: 3,
			Amount:               50,
			CreatedAt:            time.Time{},
		},
	}
	return transfers, nil
}
