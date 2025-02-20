package service

import (
	"github.com/henrriusdev/portfolio/pkg/model"
	"gorm.io/gorm"
)

type Contact struct {
	Base[model.Contact]
}

func NewContact(db *gorm.DB) *Contact {
	return &Contact{Base[model.Contact]{db}}
}
