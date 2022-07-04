package models

import (
	"learnecho/common"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100);unique_index"`
	Name     string
	Role     common.UserRole
	Password string
	// has many books
	Books []*Book `gorm:"-"`
}

func (user User) String() string {
	return user.Name
}

// func (user *User) BeforeSave() (err error) {
// 	hashed, err := common.GetPasswordUtil().HashPassword(user.Password)
// 	user.Password = hashed
// 	return
// }
