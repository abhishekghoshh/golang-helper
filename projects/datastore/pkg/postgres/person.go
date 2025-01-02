package postgres

type Person struct {
	ID       uint   `json:"id" gorm:"primary key;autoIncrement"`
	Name     string `json:"name,omitempty"`
	UserName string `json:"userName,omitempty" gorm:"unique;not null;type:varchar(25);default:null"`
	Age      int64  `json:"age,omitempty"`
	Email    string `json:"email,omitempty" gorm:"unique;not null;type:varchar(100);default:null"`
}
