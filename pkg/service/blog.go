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

func (s *Blog) GetByID(id uint) (*model.BlogPost, error) {
	var record model.BlogPost
	err := s.DB.Preload("Categories").Preload("Author").First(&record, id).Error
	return &record, err
}

func (s *Blog) GetAllWithAuthor() ([]model.BlogPost, error) {
	var posts []model.BlogPost
	err := s.DB.Preload("Author").Preload("Categories").Find(&posts).Error
	return posts, err
}

func (s *Blog) Create(post *model.BlogPost) error {
	return s.DB.Create(post).Error
}

func (s *Blog) Update(post *model.BlogPost) error {
	return s.DB.Save(post).Error
}

// SearchPosts busca posts por término de búsqueda en título y contenido
func (s *Blog) SearchPosts(query string) ([]model.BlogPost, error) {
	var posts []model.BlogPost

	// Usar LIKE para buscar en título y contenido
	err := s.DB.Preload("Author").Preload("Categories").Where(
		"title LIKE ? OR content LIKE ?",
		"%"+query+"%",
		"%"+query+"%",
	).Find(&posts).Error

	return posts, err
}

// ClearCategories elimina todas las categorías asociadas a un post
func (s *Blog) ClearCategories(postID uint) error {
	// Usar el modelo BlogPostCategory pero con el nombre de tabla correcto
	return s.DB.Exec("DELETE FROM blog_post_categories WHERE blog_post_id = ?", postID).Error
}

// GetByCategory obtiene todos los posts asociados a una categoría específica
func (s *Blog) GetByCategory(categoryID uint) ([]model.BlogPost, error) {
	var posts []model.BlogPost

	// Usar joins para obtener los posts asociados a la categoría
	err := s.DB.Joins("JOIN blog_post_categories ON blog_post_categories.blog_post_id = blog_posts.id").
		Where("blog_post_categories.category_id = ?", categoryID).
		Preload("Author").
		Preload("Categories").
		Find(&posts).Error

	return posts, err
}
