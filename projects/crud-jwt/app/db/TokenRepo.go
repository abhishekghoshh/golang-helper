package db

import (
	"crud/app/domain"
	"crud/app/error"

	"gorm.io/gorm"
)

type TokenRepo struct {
	db *gorm.DB
}

func NewTokenRepo(db *gorm.DB) *TokenRepo {
	return &TokenRepo{
		db: db,
	}
}

func (tokenRepo *TokenRepo) SaveTokenAfterLogin(loggedInUser *domain.LoggedInUser) *error.AppError {
	if err := tokenRepo.db.Create(loggedInUser).Error; err != nil {
		return error.NewUnexpectedError(err.Error())
	}
	return nil
}
func (tokenRepo *TokenRepo) GetUserEntryByUserName(username string, loggedInUser *domain.LoggedInUser) *error.AppError {
	if err := tokenRepo.db.Model(&domain.LoggedInUser{}).Where("user_name = ?", username).First(loggedInUser).Error; err != nil {
		return error.NewUnexpectedError(err.Error())
	}
	return nil
}
func (tokenRepo *TokenRepo) RemoveUserEntry(username string) *error.AppError {
	loggedInUser := domain.LoggedInUser{}
	if err := tokenRepo.GetUserEntryByUserName(username, &loggedInUser); err != nil {
		return err
	}
	if err := tokenRepo.db.Unscoped().Where("user_name = ?", username).Delete(loggedInUser).Error; err != nil {
		return error.NewUnexpectedError(err.Error())
	}
	return nil
}

func (tokenRepo *TokenRepo) RemoveUserEntryByAccessToken(accessToken string) *error.AppError {
	loggedInUser := domain.LoggedInUser{}
	if err := tokenRepo.GetUserEntryByAccessToken(accessToken, &loggedInUser); err != nil {
		return err
	}
	if err := tokenRepo.db.Unscoped().Where("access_token = ?", accessToken).Delete(loggedInUser).Error; err != nil {
		return error.NewUnexpectedError(err.Error())
	}
	return nil
}
func (tokenRepo *TokenRepo) GetUserEntryByAccessToken(accessToken string, loggedInUser *domain.LoggedInUser) *error.AppError {
	if err := tokenRepo.db.Model(&domain.LoggedInUser{}).Where("access_token = ?", accessToken).First(loggedInUser).Error; err != nil {
		return error.NewUnexpectedError(err.Error())
	}
	return nil
}
