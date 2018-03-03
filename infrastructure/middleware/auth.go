package middleware

import (
	"net/http"

	"github.com/jcox250/annadale/util/session"
)

func Authenticate(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		sessionExists, err := session.SessionExists(r, session.SessionName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if sessionExists {
			next.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, "/login/", http.StatusFound)
		}
	}
	return http.HandlerFunc(fn)
}
