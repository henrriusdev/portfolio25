package middleware

import (
	"net/http"

	"github.com/henrriusdev/portfolio/pkg/common"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := common.Store.Get(r, "session-name")
		if session.Values["authenticated"] == nil {
			http.Redirect(w, r, "/log-in", http.StatusFound)
			return
		}

		println(session.Values["authenticated"].(bool))
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Redirect(w, r, "/log-in", http.StatusFound)
			return
		}
		next(w, r)
	}
}
