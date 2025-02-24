package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-fuego/fuego"
	"github.com/henrriusdev/portfolio/api/middleware"
	"github.com/henrriusdev/portfolio/pkg/common"
	"github.com/henrriusdev/portfolio/pkg/model"
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
	fuego.Get(f, "/technologies", d.Technologies)
	fuego.Post(f, "/save-project", d.SaveProject)
	fuego.Post(f, "/save-tech", d.SaveTech)
	fuego.Delete(f, "/delete-tech/{id}", d.DeleteTech)
}

func (d *Dashboard) Index(c fuego.ContextNoBody) (fuego.Templ, error) {
	return pages.Dashboard(), nil
}

func (d *Dashboard) Projects(c fuego.ContextNoBody) (fuego.Templ, error) {
	projects, _ := d.services.Project.GetAll()
	techs, _ := d.services.Technology.GetAll()
	var baseTechs []model.Base
	for _, tech := range techs {
		baseTechs = append(baseTechs, tech)
	}
	techOptions := common.TransformToOptions(baseTechs)
	return pages.Projects(techOptions, projects...), nil
}

func (d *Dashboard) SaveProject(c fuego.ContextWithBody[model.SaveProjectRequest]) (string, error) {
	req, err := c.Body()
	if err != nil {
		return "", fuego.SendJSON(c.Response(), c.Request(), map[string]interface{}{
			"error":   "Validation failed",
			"details": err.Error(),
		})
	}

	var techs []model.Technology
	for _, idStr := range req.Technologies {
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return "", err
		}
		tech, err := d.services.Technology.GetByID(uint(id))
		if err != nil {
			return "", err
		}

		techs = append(techs, *tech)
	}

	// Create a new Project entity.
	project := model.Project{
		Title:       req.Title,
		Description: req.Description,
		ImageURL:    req.ImageURL,
		URL:         req.LiveURL,
		Repo:        req.RepoURL,
		Techs:       techs,
	}

	if err := d.services.Project.Create(&project); err != nil {
		return "", err
	}

	c.Redirect(http.StatusSeeOther, "/dashboard/projects")

	return "Project saved successfully", nil
}

func (d *Dashboard) Technologies(c fuego.ContextNoBody) (fuego.Templ, error) {
	techs, _ := d.services.Technology.GetAll()
	return pages.Techs(techs...), nil
}

func (d *Dashboard) SaveTech(c fuego.ContextWithBody[model.SaveTechRequest]) (string, error) {
	// Parse the incoming SaveTechRequest from form data.
	req, err := c.Body()
	if err != nil {
		return "", err
	}

	// Get the SVG icon URL from simpleicons.org based on the tech name.
	iconURL := common.GetSimpleIconURL(req.Name)

	// Create a new Technology entity.
	tech := model.Technology{
		Name: req.Name,
		Icon: iconURL,
	}

	if err := d.services.Technology.Create(tech); err != nil {
		return "", err
	}

	c.Redirect(http.StatusSeeOther, "/dashboard/technologies")

	// Return a success message or redirect as needed.
	return "Technology saved successfully", nil
}

func (d *Dashboard) DeleteTech(c fuego.ContextWithBody[model.DeleteTechRequest]) (string, error) {
	// Parse the incoming request to get the tech ID.
	req, err := c.Body()
	if err != nil {
		return "", err
	}
	idStr := c.PathParam("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return "", err
	}

	if req.ID == 0 {
		req.ID = uint(id)
	}

	if err := d.services.Technology.Delete(req.ID); err != nil {
		log.Println(err, "AAAAAAAAA")
		return "", err
	}

	// Return a success message or redirect as needed.
	return "Technology deleted successfully", nil
}
