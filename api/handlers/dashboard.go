package handlers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
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
	fuego.Get(f, "", d.Index,
		option.Description("Show dashboard"),
		option.Summary("Dashboard"),
		option.Tags("dashboard"),
	)
	fuego.Get(f, "/projects", d.Projects,
		option.Description("Show projects"),
		option.Summary("Projects"),
		option.Tags("dashboard", "projects"),
	)
	fuego.Get(f, "/technologies", d.Technologies,
		option.Description("Show technologies"),
		option.Summary("Technologies"),
		option.Tags("dashboard", "technologies"),
	)
	fuego.Get(f, "/experience", d.Experience,
		option.Description("Show experience"),
		option.Summary("Experience"),
		option.Tags("dashboard", "experience"),
	)
	fuego.Get(f, "/links", d.Links,
		option.Description("Show links"),
		option.Summary("Links"),
		option.Tags("dashboard", "links"),
	)
	fuego.Post(f, "/save-project", d.SaveProject,
		option.Description("Save project"),
		option.Summary("Save project"),
		option.Tags("dashboard", "projects"),
	)
	fuego.Post(f, "/save-tech", d.SaveTech,
		option.Description("Save technology"),
		option.Summary("Save technology"),
		option.Tags("dashboard", "technologies"),
	)
	fuego.Post(f, "/save-work", d.SaveWork,
		option.Description("Save work experience"),
		option.Summary("Save work experience"),
		option.Tags("dashboard", "experience"),
	)
	fuego.Delete(f, "/delete-project/{id}", d.DeleteProject,
		option.Description("Delete project"),
		option.Summary("Delete project"),
		option.Tags("dashboard", "projects"),
	)
	fuego.Delete(f, "/delete-tech/{id}", d.DeleteTech,
		option.Description("Delete technology"),
		option.Summary("Delete technology"),
		option.Tags("dashboard", "technologies"),
	)
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

func (d *Dashboard) SaveProject(c fuego.ContextWithBody[model.SaveProjectRequest]) (model.Project, error) {
	req, err := c.Body()
	if err != nil {
		return model.Project{}, fuego.SendJSON(c.Response(), c.Request(), map[string]interface{}{
			"error":   "Validation failed",
			"details": err.Error(),
		})
	}

	var techs []model.Technology
	for _, idStr := range req.Technologies {
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return model.Project{}, err
		}
		tech, err := d.services.Technology.GetByID(uint(id))
		if err != nil {
			return model.Project{}, err
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

	if req.ID != 0 {
		project.ID = req.ID
	}

	if err := d.services.Project.Upsert(&project); err != nil {
		log.Printf("Error saving project: %v", err)
		return model.Project{}, err
	}

	c.Redirect(http.StatusSeeOther, "/dashboard/projects")

	return model.Project{}, nil
}

func (d *Dashboard) SaveWork(c fuego.ContextWithBody[model.SaveWorkRequest]) (model.Experience, error) {
	req, err := c.Body()
	if err != nil {
		return model.Experience{}, fuego.SendJSON(c.Response(), c.Request(), map[string]interface{}{
			"error":   "Validation failed",
			"details": err.Error(),
		})
	}

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return model.Experience{}, err
	}

	var endDate *time.Time
	if req.EndDate != "" {
		parsedEndDate, err := time.Parse("2006-01-02", req.EndDate)
		if err != nil {
			return model.Experience{}, err
		}
		endDate = &parsedEndDate
	}

	// Create a new Experience entity.
	experience := model.Experience{
		Role:        req.Role,
		Company:     req.Company,
		Description: req.Description,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	if req.ID != 0 {
		experience.ID = req.ID
	}

	if err := d.services.Experience.Upsert(&experience); err != nil {
		log.Printf("Error saving experience: %v", err)
		return model.Experience{}, err
	}

	c.Redirect(http.StatusSeeOther, "/dashboard/experience")

	return model.Experience{}, nil
}

func (d *Dashboard) Experience(c fuego.ContextNoBody) (fuego.Templ, error) {
	experiences, _ := d.services.Experience.GetAll()
	return pages.Experience(experiences...), nil
}

func (d *Dashboard) Links(c fuego.ContextNoBody) (fuego.Templ, error) {
	return pages.Links(), nil
}

func (d *Dashboard) Technologies(c fuego.ContextNoBody) (fuego.Templ, error) {
	techs, _ := d.services.Technology.GetAll()
	return pages.Techs(techs...), nil
}

func (d *Dashboard) SaveTech(c fuego.ContextWithBody[model.SaveTechRequest]) (model.Technology, error) {
	// Parse the incoming SaveTechRequest from form data.
	req, err := c.Body()
	if err != nil {
		return model.Technology{}, err
	}

	// Get the SVG icon URL from simpleicons.org based on the tech name.
	iconURL := common.GetSimpleIconURL(req.Name)

	// Create a new Technology entity.
	tech := model.Technology{
		Name: req.Name,
		Icon: iconURL,
	}

	if err := d.services.Technology.Create(tech); err != nil {
		return model.Technology{}, err
	}

	c.Redirect(http.StatusSeeOther, "/dashboard/technologies")

	// Return a success message or redirect as needed.
	return tech, nil
}

func (d *Dashboard) DeleteTech(c fuego.ContextWithBody[model.DeleteRequest]) (model.DeleteRequest, error) {
	// Parse the incoming request to get the tech ID.
	req, err := c.Body()
	if err != nil {
		return model.DeleteRequest{}, err
	}
	idStr := c.PathParam("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return model.DeleteRequest{}, err
	}

	if req.ID == 0 {
		req.ID = uint(id)
	}

	if err := d.services.Technology.Delete(req.ID); err != nil {
		log.Println(err, "AAAAAAAAA")
		return model.DeleteRequest{}, err
	}

	// Return a success message or redirect as needed.
	return model.DeleteRequest{}, nil
}

func (d *Dashboard) DeleteProject(c fuego.ContextWithBody[model.DeleteRequest]) (string, error) {
	// Parse the incoming request to get the project ID.
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

	if err := d.services.Project.Delete(req.ID); err != nil {
		return "", err
	}

	// Return a success message or redirect as needed.
	return "Project deleted successfully", nil
}
