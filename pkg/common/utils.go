package common

import (
	"context"
	"io"
	"regexp"
	"strings"

	"github.com/a-h/templ"
	"github.com/gorilla/sessions"
	"github.com/henrriusdev/portfolio/config"
	"github.com/henrriusdev/portfolio/pkg/model"
	"github.com/henrriusdev/portfolio/src/components/selectbox"
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

func TransformToOptions(base []model.Base) []selectbox.Option {
	var options []selectbox.Option
	for _, b := range base {
		options = append(options, selectbox.Option{Label: b.GetLabel(), Value: b.GetValue()})
	}
	return options
}

func Unsafe(html string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, html)
		return
	})
}

// StripMarkdown elimina la sintaxis Markdown para mostrar texto plano
func StripMarkdown(md string) string {
	// Procesar el texto línea por línea para manejar mejor los patrones multilinea
	lines := strings.Split(md, "\n")
	for i, line := range lines {
		// Eliminar encabezados # Texto
		if strings.HasPrefix(strings.TrimSpace(line), "#") {
			line = regexp.MustCompile(`^\s*#{1,6}\s+(.+)$`).ReplaceAllString(line, "$1")
		}

		// Eliminar listas
		if matched, _ := regexp.MatchString(`^\s*[\*\-\+]\s+`, line); matched {
			line = regexp.MustCompile(`^\s*[\*\-\+]\s+(.+)$`).ReplaceAllString(line, "$1")
		}

		// Eliminar listas numeradas
		if matched, _ := regexp.MatchString(`^\s*\d+\.\s+`, line); matched {
			line = regexp.MustCompile(`^\s*\d+\.\s+(.+)$`).ReplaceAllString(line, "$1")
		}

		// Eliminar citas > texto
		if strings.HasPrefix(strings.TrimSpace(line), ">") {
			line = regexp.MustCompile(`^\s*>\s+(.+)$`).ReplaceAllString(line, "$1")
		}

		// Eliminar líneas horizontales
		if matched, _ := regexp.MatchString(`^\s*(\*\*\*|---|___)\s*$`, line); matched {
			line = ""
		}

		lines[i] = line
	}

	// Volver a unir el texto
	md = strings.Join(lines, " ")

	// Eliminar enlaces [texto](url)
	md = regexp.MustCompile(`\[([^\]]+)\]\([^\)]+\)`).ReplaceAllString(md, "$1")

	// Eliminar imágenes ![alt](url)
	md = regexp.MustCompile(`!\[[^\]]*\]\([^\)]+\)`).ReplaceAllString(md, "")

	// Eliminar énfasis *texto* o _texto_
	md = regexp.MustCompile(`\*(.*?)\*`).ReplaceAllString(md, "$1")
	md = regexp.MustCompile(`_(.*?)_`).ReplaceAllString(md, "$1")

	// Eliminar negrita **texto** o __texto__
	md = regexp.MustCompile(`\*\*(.*?)\*\*`).ReplaceAllString(md, "$1")
	md = regexp.MustCompile(`__(.*?)__`).ReplaceAllString(md, "$1")

	// Eliminar bloques de código ```código```
	md = regexp.MustCompile("```[\\s\\S]*?```").ReplaceAllString(md, "")

	// Eliminar código en línea `código`
	md = regexp.MustCompile("`([^`]+)`").ReplaceAllString(md, "$1")

	// Eliminar símbolos de lista de tareas
	md = regexp.MustCompile(`\[[ xX]\]\s+`).ReplaceAllString(md, "")

	// Eliminar múltiples espacios en blanco
	md = regexp.MustCompile(`\s+`).ReplaceAllString(md, " ")

	// Eliminar espacios en blanco al inicio y final
	return strings.TrimSpace(md)
}
