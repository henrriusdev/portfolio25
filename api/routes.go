package api

import (
	"github.com/go-fuego/fuego"
	"github.com/henrriusdev/portfolio/api/handlers"
	"github.com/henrriusdev/portfolio/pkg/service"
)

func portfolioRoutes(f *fuego.Server, services service.Service) {
	portfolio := handlers.NewPortfolio(services.User, services.Project, services.Experience, services.Technology, services.Contact)
	portfolio.RegisterRoutes(f)
}

func dashboardRoutes(f *fuego.Server, services service.Service) {
	dashboard := handlers.NewDashboard(services)
	dashboard.RegisterRoutes(fuego.Group(f, "/dashboard"))
}

func blogRoutes(f *fuego.Server, services service.Service) {
	blog := handlers.NewBlog(services)
	blog.RegisterRoutes(fuego.Group(f, "/blog"))
}

func categoryRoutes(f *fuego.Server, services service.Service) {
	category := handlers.NewCategory(services)
	categoryGroup := fuego.Group(f, "/categories")
	
	// Public routes
	fuego.Get(categoryGroup, "", category.List)
	fuego.Get(categoryGroup, "/{id}", category.View)
	
	// Protected routes - require authentication
	fuego.Post(categoryGroup, "", category.Create)
	fuego.Post(categoryGroup, "/delete", category.Delete)
	fuego.Post(categoryGroup, "/add-to-post", category.AddToPost)
	fuego.Post(categoryGroup, "/remove-from-post", category.RemoveFromPost)
}
