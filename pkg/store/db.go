package store

import (
	"fmt"
	"github.com/henrriusdev/portfolio/config"
	"github.com/henrriusdev/portfolio/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDatabase() (*gorm.DB, error) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", config.Env.DBHost, config.Env.DBUser, config.Env.DBPass, config.Env.DBName, config.Env.DBPort, config.Env.DBSSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos:", err)
	}

	// Migraciones automáticas
	if err := db.AutoMigrate(
		&model.User{},
		&model.Experience{},
		&model.Technology{},
		&model.Project{},
		&model.Contact{},
		&model.Visit{},
	); err != nil {
		log.Fatal(err)
	}

	return db, nil
}
