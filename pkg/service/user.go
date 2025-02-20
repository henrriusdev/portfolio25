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
