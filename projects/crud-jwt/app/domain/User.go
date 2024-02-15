package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string   `gorm:"column:user_name;uniqueIndex:user_name_index,sort:desc,length:20"`
	Password string   `gorm:"column:password"`
	UserType int      `gorm:"column:user_type"`
	Account  *Account `gorm:"foreignKey:UserID;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (u User) Role() string {
	if u.UserType == 0 {
		return "user"
	}
	return "admin"
}
