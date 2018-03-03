package middleware

import (
	"log"
	"net/http"
	"time"
)

type LoggerMiddleware interface {
	Log(http.Handler) http.Handler
}

type Logger struct {
}

func (l Logger) Log(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		method := r.Method
		route := r.URL.Path
		next.ServeHTTP(w, r)
		log.Printf("performed %s request to %s: took %v", method, route,
			time.Since(start))
	}
	return http.HandlerFunc(fn)
}
