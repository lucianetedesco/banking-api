package repositories

import (
	"encoding/json"
	"errors"
	"github.com/lucianetedesco/banking-api/entities"
	"gorm.io/gorm"
)

type AccountRepository interface {
	SaveAccount(account *entities.Account) (uint, error)
	GetAllAccounts() ([]entities.Account, error)
}

type AccountRepositoryDB struct {
	db *gorm.DB
}

type GormErr struct {
	Code string `json:"Code"`
}

func NewAccountRepository(DB *gorm.DB) *AccountRepositoryDB {
	return &AccountRepositoryDB{db: DB}
}

func (r *AccountRepositoryDB) SaveAccount(account *entities.Account) (uint, error) {
	err := r.db.Save(account).Error
	id := account.ID

	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError GormErr
		json.Unmarshal(byteErr, &newError)

		switch newError.Code {
		case "23505":
			return id, errors.New("CPF already exists")
		}
		return id, errors.New("internal server error")
	}

	return id, nil
}

func (r *AccountRepositoryDB) GetAllAccounts() ([]entities.Account, error) {
	var accounts []entities.Account
	err := r.db.Find(&accounts).Error
	return accounts, err
}
