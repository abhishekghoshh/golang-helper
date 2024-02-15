package dto

import (
	"crud/app/domain"
	"time"
)

type AccountDTO struct {
	UserId         string     `json:"userId,omitempty"`
	Password       string     `json:"password,omitempty"`
	UserName       string     `json:"userName,omitempty"`
	Name           string     `json:"name,omitempty"`
	Address        string     `json:"address,omitempty"`
	Phone          string     `json:"phoneNo,omitempty"`
	Email          string     `json:"email,omitempty"`
	DateOfBirth    CustomDate `json:"dateOfBirth,omitempty"`
	OpeningDate    time.Time  `json:"openingDate,omitempty"`
	LastActiveTime time.Time  `json:"lastActiveTime,omitempty"`
	Status         string     `json:"status,omitempty"`
	UserType       string     `json:"userType,omitempty"`
}

func NewAccountDTO() *AccountDTO {
	return &AccountDTO{}
}

func (accountDTO *AccountDTO) Create(account domain.Account, user domain.User) *AccountDTO {
	// accountDTO.UserId = strconv.Itoa(int(account.ID))
	accountDTO.UserName = user.UserName
	accountDTO.Name = account.Name
	accountDTO.Address = account.Address
	accountDTO.Phone = account.Phone
	accountDTO.Email = account.Email
	accountDTO.DateOfBirth = CustomDate(account.DateOfBirth)
	accountDTO.OpeningDate = account.OpeningDate
	accountDTO.LastActiveTime = account.OpeningDate
	accountDTO.Status = accountDTO.status(account)
	accountDTO.UserType = accountDTO.userType(user)
	return accountDTO
}

func (AccountDTO) status(account domain.Account) string {
	if account.Status == 0 {
		return "inactive"
	}
	return "active"
}

func (AccountDTO) userType(user domain.User) string {
	if user.UserType == 1 {
		return "admin"
	}
	return "user"
}

func (accountDTO *AccountDTO) CreateUser() *domain.User {
	return &domain.User{
		UserName: accountDTO.UserName,
		Password: accountDTO.Password,
		UserType: 0,
		Account:  accountDTO.createAccount(),
	}
}
func (accountDTO *AccountDTO) createAccount() *domain.Account {
	return &domain.Account{
		Name:           accountDTO.Name,
		Address:        accountDTO.Address,
		Phone:          accountDTO.Phone,
		Email:          accountDTO.Email,
		DateOfBirth:    time.Time(accountDTO.DateOfBirth),
		OpeningDate:    time.Now(),
		LastActiveTime: time.Now(),
		Status:         0,
	}
}
