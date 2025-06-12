package service

import (
	"github.com/henrriusdev/portfolio/pkg/model"
	"gorm.io/gorm"
)

type Blog struct {
	Base[model.BlogPost]
}

func NewBlog(db *gorm.DB) *Blog {
	return &Blog{Base[model.BlogPost]{DB: db}}
}

func (s *Blog) GetByTitle(title string) (*model.BlogPost, error) {
	var post model.BlogPost
	err := s.DB.Where("title = ?", title).First(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (s *Blog) GetWithAuthor(id uint) (*model.BlogPost, error) {
	var post model.BlogPost
	err := s.DB.Preload("Author").First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (s *Blog) GetAllWithAuthor() ([]model.BlogPost, error) {
	var posts []model.BlogPost
	err := s.DB.Preload("Author").Find(&posts).Error
	return posts, err
}

func (s *Blog) Create(post *model.BlogPost) error {
	return s.DB.Create(post).Error
}

func (s *Blog) Update(post *model.BlogPost) error {
	return s.DB.Save(post).Error
}
