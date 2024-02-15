package domain

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	UserID         uint      `gorm:"column:user_id;uniqueIndex:user_id_index,sort:desc"`
	Name           string    `gorm:"column:name"`
	Address        string    `gorm:"column:address"`
	Phone          string    `gorm:"column:phone_no;uniqueIndex:phone_no_index,sort:desc,length:20"`
	Email          string    `gorm:"column:email;uniqueIndex:email_index,sort:desc,length:50"`
	DateOfBirth    time.Time `gorm:"column:date_of_birth"`
	OpeningDate    time.Time `gorm:"column:opening_date"`
	LastActiveTime time.Time `gorm:"column:last_active_time"`
	Status         int       `gorm:"column:status"`
}
