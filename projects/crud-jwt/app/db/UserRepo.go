package db

import (
	"crud/app/domain"
	errr "crud/app/error"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}
func (userRepo *UserRepo) CreateUser(user *domain.User) *errr.AppError {
	tx := userRepo.db.Begin()
	defer func(tx *gorm.DB) {
		if r := recover(); r != nil && tx != nil {
			tx.Rollback()
		}
	}(tx)
	if err := tx.Error; err != nil {
		return errr.NewUnexpectedError(err.Error())
	}
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.Omit(clause.Associations).Create(user).Error; err != nil {
			return err
		}
		user.Account.UserID = user.ID
		if err := tx.Create(user.Account).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return errr.NewUnexpectedError(err.Error())
	}
	if err := tx.Commit().Error; err != nil {
		return errr.NewUnexpectedError(err.Error())
	}
	return nil
}
func (userRepo *UserRepo) GetUserById(user *domain.User, userId uint) *errr.AppError {
	err := userRepo.db.Model(&domain.User{}).Preload("Account").Where("id = ?", userId).First(user).Error
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return errr.NewNotFoundError("user not found")
		}
		return errr.NewUnexpectedError(err.Error())
	}
	return nil
}
func (userRepo *UserRepo) GetUserByUserName(user *domain.User, username string) *errr.AppError {
	err := userRepo.db.Model(&domain.User{}).Where("user_name = ?", username).First(user).Error
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return errr.NewNotFoundError("user not found")
		}
		return errr.NewUnexpectedError(err.Error())
	}
	return nil
}
func (userRepo *UserRepo) GetUserWithAccountByUserName(user *domain.User, username string) *errr.AppError {
	err := userRepo.db.Model(&domain.User{}).Preload("Account").Where("user_name = ?", username).First(user).Error
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return errr.NewNotFoundError("user not found")
		}
		return errr.NewUnexpectedError(err.Error())
	}
	return nil
}
func (userRepo *UserRepo) GetUsers(users *[]domain.User) *errr.AppError {
	if err := userRepo.db.Model(&domain.User{}).Preload("Account").Find(users).Error; err != nil {
		return errr.NewUnexpectedError(err.Error())
	}
	return nil
}

func (userRepo *UserRepo) UpdateUser(user *domain.User) *errr.AppError {
	err := userRepo.db.Save(user).Error
	if err != nil {
		return errr.NewUnexpectedError(err.Error())
	}
	return nil
}

func (userRepo *UserRepo) DeleteUser(id uint) *errr.AppError {
	var user domain.User
	if err := userRepo.GetUserById(&user, id); err != nil {
		return err
	}
	tx := userRepo.db.Begin()
	defer func(tx *gorm.DB) {
		if r := recover(); r != nil && tx != nil {
			tx.Rollback()
		}
	}(tx)
	if err := tx.Error; err != nil {
		return errr.NewUnexpectedError(err.Error())
	}
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Where("id = ?", id).Delete(&user).Error; err != nil {
			return err
		}
		if err := tx.Unscoped().Where("id = ?", user.Account.ID).Delete(user.Account).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return errr.NewUnexpectedError(err.Error())
	}
	if err := tx.Commit().Error; err != nil {
		return errr.NewUnexpectedError(err.Error())
	}
	return nil
}
