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

// GetAll returns all experiences
func (e *Experience) GetAll() ([]model.Experience, error) {
	var experiences []model.Experience
	err := e.DB.Order("start_date desc").Find(&experiences).Error
	return experiences, err
}
