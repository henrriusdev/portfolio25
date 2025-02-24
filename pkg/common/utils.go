package common

import (
	"strings"

	"github.com/gorilla/sessions"
	"github.com/henrriusdev/portfolio/config"
	"github.com/henrriusdev/portfolio/pkg/model"
	"github.com/henrriusdev/portfolio/src/components"
	"golang.org/x/crypto/bcrypt"
)

var Store = sessions.NewCookieStore([]byte(config.Env.JWTSecret))

func InitStore() {
	Store = sessions.NewCookieStore([]byte(config.Env.JWTSecret))
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPasswordHash compares a plain text password with a hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetSimpleIconURL(iconName string) string {
	if strings.Contains(iconName, " ") {
		iconName = strings.ReplaceAll(iconName, " ", "")
	}

	if strings.Contains(iconName, ".") {
		iconName = strings.ReplaceAll(iconName, ".", "dot")
	}
	return "https://cdn.simpleicons.org/" + strings.ToLower(iconName)
}

func TransformToOptions(base []model.Base) []components.SelectOption {
	var options []components.SelectOption
	for _, b := range base {
		options = append(options, components.SelectOption{Label: b.GetLabel(), Value: b.GetValue()})
	}
	return options
}
