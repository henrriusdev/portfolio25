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

func (p *Project) GetAll() ([]model.Project, error) {
	var projects []model.Project
	err := p.DB.Find(&projects).Error
	return projects, err
}
