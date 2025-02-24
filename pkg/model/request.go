package model

type LoginRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type SaveTechRequest struct {
	Name string `form:"name"` // Technology name
}

type DeleteTechRequest struct {
	ID uint `param:"id"` // Technology ID
}

type SaveProjectRequest struct {
	Title        string   `schema:"title"`
	Description  string   `schema:"description"`
	Technologies []string `schema:"techs[]"`
	ImageURL     string   `schema:"image_url"`
	LiveURL      string   `schema:"url"`
	RepoURL      string   `schema:"repo"`
}
