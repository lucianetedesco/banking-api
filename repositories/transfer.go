package repositories

import (
	"github.com/lucianetedesco/banking-api/entities"
	"gorm.io/gorm"
)

type TransferRepository interface {
	SaveTransfer(transfer *entities.Transfer) (uint, error)
	GetAllTransfers(accountId string) ([]entities.Transfer, error)
}

type TransferRepositoryDB struct {
	db *gorm.DB
}

func NewTransferRepository(DB *gorm.DB) *TransferRepositoryDB {
	return &TransferRepositoryDB{db: DB}
}

func (r *TransferRepositoryDB) SaveTransfer(transfer *entities.Transfer) (uint, error) {
	return transfer.ID, r.db.Save(transfer).Error
}

func (r *TransferRepositoryDB) GetAllTransfers(accountId string) ([]entities.Transfer, error) {
	var transfers []entities.Transfer
	err := r.db.Where("account_origin_id = ?", accountId).First(&transfers).Error
	return transfers, err
}
