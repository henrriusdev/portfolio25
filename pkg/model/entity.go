package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

type Experience struct {
	gorm.Model
	Company     string    `gorm:"not null"`
	Role        string    `gorm:"not null"`
	StartDate   time.Time `gorm:"not null"`
	EndDate     *time.Time
	Description string
}

type Technology struct {
	gorm.Model
	Name string `gorm:"not null"`
	Icon string `gorm:"not null"`
}

type Project struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	URL         string
	Repo        string
}

type Contact struct {
	gorm.Model
	Platform string `gorm:"not null"`
	URL      string `gorm:"not null"`
}

type Visit struct {
	gorm.Model
	Page      string    `gorm:"not null"`
	VisitDate time.Time `gorm:"not null;default:CURRENT_DATE"`
	Count     int       `gorm:"default:1"`
}
