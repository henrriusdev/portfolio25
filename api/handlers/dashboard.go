package handlers

import (
	"net/http"

	"github.com/go-fuego/fuego"
	"github.com/henrriusdev/portfolio/api/middleware"
	"github.com/henrriusdev/portfolio/pkg/service"
	"github.com/henrriusdev/portfolio/src/pages"
)

type Dashboard struct {
	services service.Service
}

func NewDashboard(services service.Service) *Dashboard {
	return &Dashboard{services}
}

func (d *Dashboard) RegisterRoutes(f *fuego.Server) {
	fuego.Use(f, func(next http.Handler) http.Handler {
		return middleware.AuthMiddleware(http.HandlerFunc(next.ServeHTTP))
	})
	fuego.Get(f, "", d.Index)
	fuego.Get(f, "/projects", d.Projects)
}

func (d *Dashboard) Index(c fuego.ContextNoBody) (fuego.Templ, error) {
	return pages.Dashboard(), nil
}

func (d *Dashboard) Projects(c fuego.ContextNoBody) (fuego.Templ, error) {
	projects, _ := d.services.Project.GetAll()
	return pages.Projects(projects...), nil
}
