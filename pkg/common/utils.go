package common

import (
	"github.com/gorilla/sessions"
	"github.com/henrriusdev/portfolio/config"
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
