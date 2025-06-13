package service

import (
	"github.com/henrriusdev/portfolio/pkg/model"
	"gorm.io/gorm"
)

type Category struct {
	db *gorm.DB
}

func NewCategory(db *gorm.DB) *Category {
	return &Category{db: db}
}

// GetAll retrieves all categories
func (c *Category) GetAll() ([]model.Category, error) {
	var categories []model.Category
	if err := c.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

// GetByID retrieves a category by its ID
func (c *Category) GetByID(id uint) (*model.Category, error) {
	var category model.Category
	if err := c.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

// Create creates a new category
func (c *Category) Create(category *model.Category) error {
	return c.db.Create(category).Error
}

// Update updates an existing category
func (c *Category) Update(category *model.Category) error {
	return c.db.Save(category).Error
}

// Delete deletes a category by its ID
func (c *Category) Delete(id uint) error {
	return c.db.Delete(&model.Category{}, id).Error
}

// GetPostCategories retrieves all categories for a specific post
func (c *Category) GetPostCategories(postID uint) ([]model.Category, error) {
	var categories []model.Category
	if err := c.db.Joins("JOIN blog_post_categories ON blog_post_categories.category_id = categories.id").
		Where("blog_post_categories.blog_post_id = ?", postID).
		Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

// AddCategoryToPost adds a category to a post
func (c *Category) AddCategoryToPost(postID, categoryID uint) error {
	return c.db.Create(&model.BlogPostCategory{
		BlogPostID: postID,
		CategoryID: categoryID,
	}).Error
}

// RemoveCategoryFromPost removes a category from a post
func (c *Category) RemoveCategoryFromPost(postID, categoryID uint) error {
	return c.db.Where("blog_post_id = ? AND category_id = ?", postID, categoryID).
		Delete(&model.BlogPostCategory{}).Error
}
