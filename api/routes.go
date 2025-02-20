package api

import (
	"github.com/go-fuego/fuego"
	"github.com/henrriusdev/portfolio/api/handlers"
	"github.com/henrriusdev/portfolio/pkg/service"
)

func portfolioRoutes(f *fuego.Server, services service.Service) {
	portfolio := handlers.NewPortfolio(services.User)
	portfolio.RegisterRoutes(f)
}
