package service

import (
	"github.com/henrriusdev/portfolio/pkg/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Project struct {
	Base[model.Project]
}

func NewProject(db *gorm.DB) *Project {
	return &Project{Base[model.Project]{db}}
}

func (p *Project) GetAll() ([]model.Project, error) {
	var projects []model.Project
	err := p.DB.Preload("Techs").Find(&projects).Error
	return projects, err
}

func (p *Project) Upsert(project *model.Project) error {
	return p.DB.Clauses(clause.OnConflict{UpdateAll: true}).Create(project).Error
}

func (p *Project) Delete(id uint) error {
	return p.DB.Delete(&model.Project{}, id).Error
}
