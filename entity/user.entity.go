package entity

import (
	"github.com/mr687/simple-go-rest-api/helper"
)

// User struct is represents a users table in database
type User struct {
	Id       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Username string `gorm:"uniqueIndex;type:varchar(255)" json:"username"`
	Email    string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string `gorm:"->;<-;no null" json:"-"`
}

func (u *User) BeforeInsert() error {
	hashedPassword, err := helper.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	return nil
}
