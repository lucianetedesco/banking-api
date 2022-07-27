package entities

import (
	"errors"
	"testing"
)

func TestTransferValidateSuccess(t *testing.T) {
	transfer := Transfer{
		AccountDestinationId: 1542,
		Amount:               100,
	}
	result := transfer.Validate()

	if result != nil {
		t.Errorf("result %v, expected %v", result, nil)
	}
}

func TestTransferValidateInvalidAmount(t *testing.T) {
	transfer := Transfer{
		AccountDestinationId: 1542,
		Amount:               0,
	}
	result := transfer.Validate()
	expected := errors.New("mount must be greater than 0")

	if result.Error() != expected.Error() {
		t.Errorf("result %v, expected %v", result, expected)
	}
}
