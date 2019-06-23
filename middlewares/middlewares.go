package middlewares

import (
	"log"
	"net/http"
	"os"
	"time"
)

type statusWriter struct {
	http.ResponseWriter
	status      int
	written     int64
	wroteHeader bool
}

func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Api-Key") == os.Getenv("API_KEY") {
			h.ServeHTTP(w, r)
			return

		}
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("403 - Invalid Api Key"))
		return
	})
}

func LoggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		sw := statusWriter{ResponseWriter: w}
		h.ServeHTTP(&sw, r)
		log.Printf("%s - %s - %s (%v)\n",
					r.Method, r.URL.Path,
					r.RequestURI,
					time.Since(startTime))
		return

	})

}
