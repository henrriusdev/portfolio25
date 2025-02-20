package assets

import (
	"embed"
	"io"
	"net/http"
	"strings"
)

//go:embed *
var StaticFiles embed.FS

func Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/static/")

		file, err := StaticFiles.Open(path)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		defer file.Close()

		content, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, "500 - Error interno del servidor", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", http.DetectContentType(content))
		w.Write(content)
	})
}
