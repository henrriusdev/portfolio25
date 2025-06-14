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
	err := p.DB.Preload("Techs").Find(&projects).Error
	return projects, err
}

func (p *Project) Upsert(project *model.Project) error {
	// Si es una actualización (ID existente)
	if project.ID != 0 {
		// 1. Obtener el proyecto existente con sus tecnologías
		existingProject := model.Project{}
		if err := p.DB.Preload("Techs").First(&existingProject, project.ID).Error; err != nil {
			return err
		}

		// 2. Guardar una copia de las tecnologías actuales
		newTechs := project.Techs

		// 3. Actualizar los campos básicos del proyecto
		project.CreatedAt = existingProject.CreatedAt // Preservar fecha de creación
		if err := p.DB.Model(project).Updates(map[string]interface{}{
			"title":       project.Title,
			"description": project.Description,
			"image_url":   project.ImageURL,
			"url":         project.URL,
			"repo":        project.Repo,
		}).Error; err != nil {
			return err
		}

		// 4. Eliminar todas las relaciones existentes
		if err := p.DB.Model(&existingProject).Association("Techs").Clear(); err != nil {
			return err
		}

		// 5. Establecer las nuevas relaciones
		return p.DB.Model(&existingProject).Association("Techs").Replace(newTechs)
	}

	// Si es una creación (ID nuevo)
	return p.DB.Create(project).Error
}

func (p *Project) Delete(id uint) error {
	return p.DB.Delete(&model.Project{}, id).Error
}
