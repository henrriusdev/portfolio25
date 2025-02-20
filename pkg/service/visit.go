package service

import (
	"github.com/henrriusdev/portfolio/pkg/model"
	"gorm.io/gorm"
)

type Visit struct {
	Base[model.Visit]
}

func NewVisit(db *gorm.DB) *Visit {
	return &Visit{Base[model.Visit]{db}}
}
