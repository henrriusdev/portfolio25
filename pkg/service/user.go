package service

import (
	"github.com/henrriusdev/portfolio/pkg/model"
	"gorm.io/gorm"
)

type User struct {
	Base[model.User]
}

func NewUser(db *gorm.DB) *User {
	return &User{Base[model.User]{db}}
}

func (u *User) Login(username, password string) error {
	var user model.User
	result := u.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
