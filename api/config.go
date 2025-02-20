package api

import (
	"github.com/go-fuego/fuego"
	"github.com/henrriusdev/portfolio/pkg/service"
	"github.com/henrriusdev/portfolio/src/assets"
	"gorm.io/gorm"
	"log"
)

func Start(db *gorm.DB) {
	s := fuego.NewServer()
	//fuego.Handle(s, "/static/{filepath...}", assets.Handler())
	fuego.Handle(s, "/static/", assets.Static())
	services := newServices(db)

	portfolioRoutes(s, services)
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
		Visit:      service.NewVisit(db),
	}
}
