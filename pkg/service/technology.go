package service

import (
	"github.com/henrriusdev/portfolio/pkg/model"
	"gorm.io/gorm"
)

type Technology struct {
	Base[model.Technology]
}

func NewTechnology(db *gorm.DB) *Technology {
	return &Technology{Base[model.Technology]{db}}
}
