package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-fuego/fuego"
	"github.com/henrriusdev/portfolio/pkg/model"
	"github.com/henrriusdev/portfolio/pkg/service"
	"github.com/henrriusdev/portfolio/src/pages"
)

type Category struct {
	services service.Service
}

func NewCategory(services service.Service) *Category {
	return &Category{services: services}
}

// List returns all categories
func (c *Category) List(ctx fuego.ContextNoBody) (fuego.Templ, error) {
	categories, err := c.services.Category.GetAll()
	if err != nil {
		return nil, err
	}

	return pages.CategoryList(categories), nil
}

// Create creates a new category
func (c *Category) Create(ctx fuego.ContextWithBody[model.Category]) (fuego.Templ, error) {
	category, err := ctx.Body()
	if err != nil {
		return nil, err
	}
	err = c.services.Category.Create(&category)
	if err != nil {
		return nil, err
	}

	// Redirect to the category list page
	ctx.Redirect(http.StatusSeeOther, "/categories")
	return nil, nil
}

// Delete deletes a category
type DeleteCategoryRequest struct {
	ID uint `form:"id"`
}

func (c *Category) Delete(ctx fuego.ContextWithBody[DeleteCategoryRequest]) (fuego.Templ, error) {
	request, err := ctx.Body()
	if err != nil {
		return nil, err
	}

	err = c.services.Category.Delete(request.ID)
	if err != nil {
		return nil, err
	}

	// Redirect to the category list page
	ctx.Redirect(http.StatusSeeOther, "/categories")
	return nil, nil
}

// AddToPost adds a category to a blog post
type CategoryPostRequest struct {
	PostID     uint `form:"post_id"`
	CategoryID uint `form:"category_id"`
}

func (c *Category) AddToPost(ctx fuego.ContextWithBody[CategoryPostRequest]) (fuego.Templ, error) {
	request, err := ctx.Body()
	if err != nil {
		return nil, err
	}

	err = c.services.Category.AddCategoryToPost(request.PostID, request.CategoryID)
	if err != nil {
		return nil, err
	}
	
	// Redirect back to the blog edit page
	ctx.Redirect(http.StatusSeeOther, fmt.Sprintf("/blog/edit/%d", request.PostID))
	return nil, nil
}

// RemoveFromPost removes a category from a blog post
func (c *Category) RemoveFromPost(ctx fuego.ContextWithBody[CategoryPostRequest]) (fuego.Templ, error) {
	request, err := ctx.Body()
	if err != nil {
		return nil, err
	}

	err = c.services.Category.RemoveCategoryFromPost(request.PostID, request.CategoryID)
	if err != nil {
		return nil, err
	}
	
	// Redirect back to the blog edit page
	ctx.Redirect(http.StatusSeeOther, fmt.Sprintf("/blog/edit/%d", request.PostID))
	return nil, nil
}

// View shows posts associated with a specific category
func (c *Category) View(ctx fuego.ContextNoBody) (fuego.Templ, error) {
	// Get category ID from URL parameter
	categoryIDStr := ctx.PathParam("id")
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
	if err != nil {
		return nil, err
	}

	// Get category details
	category, err := c.services.Category.GetByID(uint(categoryID))
	if err != nil {
		return nil, err
	}

	// Get posts for this category
	posts, err := c.services.Blog.GetByCategory(uint(categoryID))
	if err != nil {
		return nil, err
	}

	// Return the category view template
	return pages.CategoryView(*category, posts), nil
}
