package handlers

import (
	"log"
	"net/http"

	"github.com/go-fuego/fuego"
	"github.com/henrriusdev/portfolio/pkg/common"
	"github.com/henrriusdev/portfolio/pkg/model"
	"github.com/henrriusdev/portfolio/pkg/service"
	"github.com/henrriusdev/portfolio/src/pages"
)

type Portfolio struct {
	service *service.User
	project *service.Project
	exp     *service.Experience
	contact *service.Contact
	tech    *service.Technology
}

func NewPortfolio(user *service.User, project *service.Project, exp *service.Experience, tech *service.Technology, contact *service.Contact) *Portfolio {
	return &Portfolio{user, project, exp, contact, tech}
}

func (p *Portfolio) RegisterRoutes(f *fuego.Server) {
	fuego.All(f, "/", p.Index)
	fuego.Get(f, "/log-in", p.Login)
	fuego.Post(f, "/log-in", p.LoginPost)
}

func (p *Portfolio) Index(c fuego.ContextNoBody) (fuego.Templ, error) {
	projects, _ := p.project.GetAll()
	experience, _ := p.exp.GetAll()
	technologies, _ := p.tech.GetAll()
	links, _ := p.contact.GetAll()
	return pages.HomePage(projects, experience, technologies, links), nil
}

func (p *Portfolio) Login(c fuego.ContextNoBody) (fuego.Templ, error) {
	return pages.Login(), nil
}

func (p *Portfolio) LoginPost(c fuego.ContextWithBody[model.LoginRequest]) (model.Response, error) {
	body, err := c.Body()
	if err != nil {
		return model.Response{Success: false}, err
	}

	if err := p.service.Login(body.Username, body.Password); err != nil {
		return model.Response{Success: false}, err
	}

	session, _ := common.Store.Get(c.Request(), "session-name")
	session.Values["authenticated"] = true
	if err := session.Save(c.Request(), c.Response()); err != nil {
		log.Println(err)
		return model.Response{Success: false}, err
	}

	c.Redirect(http.StatusSeeOther, "/dashboard")

	return model.Response{Success: true}, nil
}
