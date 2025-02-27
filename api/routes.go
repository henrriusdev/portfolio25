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
