package model

import (
	"context"
	"fmt"

	"github.com/go-fuego/fuego"
)

type LoginRequest struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

type SaveTechRequest struct {
	Name string `form:"name" json:"name"` // Technology name
}

type DeleteRequest struct {
	ID uint `param:"id"` // Technology ID
}

type SaveProjectRequest struct {
	ID           uint     `schema:"id,omitempty" json:"id,omitempty"`
	Title        string   `schema:"title" json:"title,omitempty"`
	Description  string   `schema:"description" json:"description,omitempty"`
	Technologies []string `schema:"techs" json:"techs,omitempty"`
	ImageURL     string   `schema:"image_url" json:"image_url,omitempty"`
	LiveURL      string   `schema:"url" json:"url,omitempty"`
	RepoURL      string   `schema:"repo" json:"repo,omitempty"`
}

type SaveWorkRequest struct {
	ID          uint   `schema:"id,omitempty" json:"id,omitempty"`
	Role        string `schema:"role" json:"role,omitempty"`
	Company     string `schema:"company" json:"company,omitempty"`
	Description string `schema:"description" json:"description,omitempty"`
	StartDate   string `schema:"start_date" json:"start_date,omitempty"`
	EndDate     string `schema:"end_date" json:"end_date,omitempty"`
	IsCurrent   bool   `schema:"is_current" json:"is_current,omitempty"`
}

type SaveLinkRequest struct {
	ID       uint   `schema:"id,omitempty" json:"id,omitempty"`
	Platform string `schema:"platform" json:"platform"`
	URL      string `schema:"url" json:"url"`
}

func (r *SaveWorkRequest) InTransform(ctx context.Context) error {
	// Handle start_date: trim if too long or ensure it's not empty
	if r.StartDate == "" {
		return fmt.Errorf("start date cannot be empty")
	} else if len(r.StartDate) > 10 {
		r.StartDate = r.StartDate[:10]
	}

	// Handle end_date: if is_current is true, ensure end_date is empty
	// El checkbox is_current enviará "true" como string cuando está marcado
	if r.IsCurrent {
		// Si es trabajo actual, la fecha de fin debe estar vacía
		r.EndDate = ""
	} else if len(r.EndDate) > 10 {
		// Si no es trabajo actual pero tiene fecha de fin, la recortamos si es necesario
		r.EndDate = r.EndDate[:10]
	}

	return nil
}

var _ fuego.InTransformer = (*SaveWorkRequest)(nil)
