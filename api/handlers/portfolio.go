package handlers

import (
	"github.com/go-fuego/fuego"
	"github.com/henrriusdev/portfolio/pkg/service"
	"github.com/henrriusdev/portfolio/src/pages"
)

type Portfolio struct {
	service *service.User
}

func NewPortfolio(user *service.User) *Portfolio {
	return &Portfolio{user}
}

func (p *Portfolio) RegisterRoutes(f *fuego.Server) {
	fuego.Get(f, "/", p.Index)
}

func (p *Portfolio) Index(c fuego.ContextNoBody) (fuego.Gomponent, error) {
	return pages.HomePage(), nil
}
