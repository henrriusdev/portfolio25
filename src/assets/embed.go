package assets

import (
	"embed"
	"net/http"
)

//go:embed *
var StaticFiles embed.FS

func Static() http.Handler {
	return http.StripPrefix("/static/", http.FileServer(http.FS(StaticFiles)))
}
