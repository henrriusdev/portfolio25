package api

import (
	"log"

	"github.com/go-fuego/fuego"
	"github.com/henrriusdev/portfolio/pkg/common"
	"github.com/henrriusdev/portfolio/pkg/service"
	"github.com/henrriusdev/portfolio/src/assets"
	"gorm.io/gorm"
)

func Start(db *gorm.DB) {
	s := fuego.NewServer()
	s.OpenAPI.Description().Info.Title = "Portfolio API"
	s.OpenAPI.Description().Info.Version = "1.0.0"
	common.InitStore()
	// fuego.Handle(s, "/static/{filepath...}", assets.Handler())
	fuego.Handle(s, "/static/", assets.Static())
	services := newServices(db)

	portfolioRoutes(s, services)
	dashboardRoutes(s, services)
	blogRoutes(s, services)
	categoryRoutes(s, services)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}

func newServices(db *gorm.DB) service.Service {
	return service.Service{
		Contact:    service.NewContact(db),
		Experience: service.NewExperience(db),
		Project:    service.NewProject(db),
		Technology: service.NewTechnology(db),
		User:       service.NewUser(db),
		Blog:       service.NewBlog(db),
		Visit:      service.NewVisit(db),
		Category:   service.NewCategory(db),
	}
}
