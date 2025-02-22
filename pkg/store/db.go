package store

import (
	"fmt"
	"log"

	"github.com/henrriusdev/portfolio/config"
	"github.com/henrriusdev/portfolio/pkg/common"
	"github.com/henrriusdev/portfolio/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
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
		&model.BlogPost{},
		&model.Category{},
	); err != nil {
		log.Fatal(err)
	}

	// Create default user if it doesn't exist
	var defaultUser model.User
	result := db.Where("username = ?", config.Env.DefaultUsername).First(&defaultUser)
	if result.Error == gorm.ErrRecordNotFound {
		hashedPassword, err := common.HashPassword(config.Env.DefaultPassword)
		if err != nil {
			log.Printf("Error hashing password: %v", err)
			return db, err
		}

		defaultUser = model.User{
			Username: config.Env.DefaultUsername,
			Password: hashedPassword,
		}
		if err := db.Create(&defaultUser).Error; err != nil {
			log.Printf("Error creating default user: %v", err)
			return db, err
		}
		log.Printf("Default user created successfully")
	}

	return db, nil
}
