package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Base interface {
	GetLabel() string
	GetValue() string
}

type User struct {
	Username string     `gorm:"unique;not null"`
	Password string     `gorm:"not null"`
	Posts    []BlogPost `gorm:"foreignKey:AuthorID"`
	gorm.Model
}

type Experience struct {
	Company     string    `gorm:"not null"`
	Role        string    `gorm:"not null"`
	StartDate   time.Time `gorm:"not null"`
	EndDate     *time.Time
	Description string
	gorm.Model
}

type Technology struct {
	Name string `json:"name" gorm:"not null"`
	Icon string `json:"icon" gorm:"not null"`
	gorm.Model
}

func (t Technology) GetLabel() string {
	return t.Name
}

func (t Technology) GetValue() string {
	return fmt.Sprintf("%d", t.ID)
}

type Project struct {
	Title       string       `json:"title" gorm:"not null"`
	Description string       `json:"description" gorm:"not null"`
	ImageURL    string       `json:"imageUrl" gorm:"not null"`
	URL         string       `json:"url"`
	Repo        string       `json:"repo"`
	Techs       []Technology `json:"techs" gorm:"many2many:project_technologies;"`
	gorm.Model
}

func (p Project) GetLabel() string {
	return p.Title
}

func (p Project) GetValue() string {
	return fmt.Sprintf("%d", p.ID)
}

type Contact struct {
	Platform string `gorm:"not null"`
	Icon     string `gorm:"not null"`
	URL      string `gorm:"not null"`
	gorm.Model
}

type Visit struct {
	Page      string    `gorm:"not null"`
	VisitDate time.Time `gorm:"not null;default:CURRENT_DATE"`
	Count     int       `gorm:"default:1"`
	gorm.Model
}

type BlogPost struct {
	Title      string     `gorm:"not null"`
	Content    string     `gorm:"type:text;not null"` // Markdown content
	ImageURL   string     // Optional field for storing image URL
	AuthorID   uint       `gorm:"not null"`
	Author     User       // Foreign key relationship with User
	Categories []Category `gorm:"many2many:blog_post_categories;"` // Many-to-many relationship
	gorm.Model
}

type Category struct {
	Name        string `gorm:"not null"`
	Description string
	BlogPosts   []BlogPost `gorm:"many2many:blog_post_categories;"` // Many-to-many relationship
	gorm.Model
}

type BlogPostCategory struct {
	BlogPostID uint `gorm:"primaryKey"`
	CategoryID uint `gorm:"primaryKey"`
}
