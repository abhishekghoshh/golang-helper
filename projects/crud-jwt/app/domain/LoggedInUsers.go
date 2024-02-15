package domain

import "gorm.io/gorm"

type LoggedInUser struct {
	gorm.Model
	UserName     string `gorm:"column:user_name;uniqueIndex:user_name_index,sort:desc,length:20"`
	AccessToken  string `gorm:"column:access_token"`
	RefreshToken string `gorm:"column:refresh_token"`
}

func NewLoggedInUser(username string, accessToken string, refreshToken string) *LoggedInUser {
	return &LoggedInUser{
		UserName:     username,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
