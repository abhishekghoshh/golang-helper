package db

import (
	"crud/app/domain"
	"crud/app/error"

	"gorm.io/gorm"
)

type AccountRepo struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) *AccountRepo {
	return &AccountRepo{
		db: db,
	}
}
func (accountRepo *AccountRepo) GetAll() ([]domain.Account, *error.AppError) {
	accounts := make([]domain.Account, 0)

	return accounts, nil
}

func (accountRepo *AccountRepo) GetById(userId int32) (domain.Account, *error.AppError) {
	var account domain.Account

	return account, nil
}

func (accountRepo *AccountRepo) AddAccount(account *domain.Account) *error.AppError {
	err := accountRepo.db.Create(account).Error
	if err != nil {
		return error.NewUnexpectedError(err.Error())
	}
	return nil
}
