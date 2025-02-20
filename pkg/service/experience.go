package service

import (
	"github.com/henrriusdev/portfolio/pkg/model"
	"gorm.io/gorm"
)

type Experience struct {
	Base[model.Experience]
}

func NewExperience(db *gorm.DB) *Experience {
	return &Experience{Base[model.Experience]{db}}
}
