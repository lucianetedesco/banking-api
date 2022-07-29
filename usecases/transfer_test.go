package usecases

import (
	"github.com/lucianetedesco/banking-api/entities"
	"github.com/lucianetedesco/banking-api/settings"
	"reflect"
	"testing"
	"time"
)

func TestGetAllTransfer(t *testing.T) {
	settings.InitConfig()
	expected := []entities.Transfer{
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

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjo1MiwiZXhwIjoxNjU5MDYxNzQ1fQ.HF2Trjtz5hfZVi7NOXGNSa0OyEx8XPCT4waRpqYBfIo"
	useCaseTransfer := NewTransferUseCase(TransferRepositoryMock{})
	result, err := useCaseTransfer.GetTransfers(token)

	if !reflect.DeepEqual(result, expected) && err != nil {
		t.Errorf("result %v, expected %v", result, expected)
	}
}
