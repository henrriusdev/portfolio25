package service

import (
	"github.com/henrriusdev/portfolio/pkg/model"
	"gorm.io/gorm"
)

type Project struct {
	Base[model.Project]
}

func NewProject(db *gorm.DB) *Project {
	return &Project{Base[model.Project]{db}}
}
