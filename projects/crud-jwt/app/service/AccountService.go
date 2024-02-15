package service

import (
	"crud/app/db"
	"crud/app/domain"
	"crud/app/dto"
	"crud/app/error"
	"crud/app/logger"
	"crud/app/util"
	"strconv"
)

type AccountService struct {
	accountRepo *db.AccountRepo
	userRepo    *db.UserRepo
	encoder     *util.PasswordEncoder
}

func (accountService *AccountService) GetAccounts() ([]dto.AccountDTO, *error.AppError) {
	users := make([]domain.User, 0)
	if err := accountService.userRepo.GetUsers(&users); err != nil {
		return nil, err
	}
	accountDtos := make([]dto.AccountDTO, 0)
	for _, user := range users {
		var accountDto dto.AccountDTO
		accountDtos = append(accountDtos, *accountDto.Create(*user.Account, user))
	}
	return accountDtos, nil
}
func (accountService *AccountService) GetAccount(pathVariables map[string]string, accountDto *dto.AccountDTO) *error.AppError {
	accountId, _ := strconv.Atoi(pathVariables["accountId"])
	user := domain.User{}
	if err := accountService.userRepo.GetUserById(&user, uint(accountId)); err != nil {
		return err
	}
	accountDto.Create(*user.Account, user)
	return nil
}
func (accountService *AccountService) AddAccount(accountDto *dto.AccountDTO, accountOperation *dto.AccountOperation) *error.AppError {
	var err *error.AppError
	accountDto.Password, err = accountService.encoder.Encrypt(accountDto.Password)
	if err != nil {
		return err
	}
	user := accountDto.CreateUser()
	err = accountService.userRepo.CreateUser(user)
	if err != nil {
		return err
	}
	accountOperation.CreateAccountOperation(strconv.Itoa(int(user.ID)), "user created")
	logger.Info("user created with user id " + accountOperation.UserId)
	return nil
}
func (accountService *AccountService) DeleteAccount(pathVariables map[string]string, accountOperation *dto.AccountOperation) *error.AppError {
	accountId, _ := strconv.Atoi(pathVariables["accountId"])
	if err := accountService.userRepo.DeleteUser(uint(accountId)); err != nil {
		return err
	}
	accountOperation.CreateAccountOperation(pathVariables["accountId"], "user deleted")
	logger.Info("user deleted with user id " + accountOperation.UserId)
	return nil
}
func (accountService *AccountService) UpdateAccount(accountDto *dto.AccountDTO, accountOperation *dto.AccountOperation) *error.AppError {
	return nil
}

func NewAccountService(accountRepo *db.AccountRepo, userRepo *db.UserRepo, encoder *util.PasswordEncoder) *AccountService {
	return &AccountService{
		accountRepo: accountRepo,
		userRepo:    userRepo,
		encoder:     encoder,
	}
}
