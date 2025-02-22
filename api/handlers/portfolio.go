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
	fuego.Get(f, "/home", p.Index)
	fuego.Get(f, "/log-in", p.Login)
}

func (p *Portfolio) Index(c fuego.ContextNoBody) (fuego.Templ, error) {
	return pages.HomePage(), nil
}

func (p *Portfolio) Login(c fuego.ContextNoBody) (fuego.Templ, error) {
	return pages.Login(), nil
}
