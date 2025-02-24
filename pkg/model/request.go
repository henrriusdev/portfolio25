package model

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
	Technologies []string `schema:"techs[]" json:"techs,omitempty"`
	ImageURL     string   `schema:"image_url" json:"image_url,omitempty"`
	LiveURL      string   `schema:"url" json:"url,omitempty"`
	RepoURL      string   `schema:"repo" json:"repo,omitempty"`
}
