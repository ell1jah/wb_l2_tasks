package handler

import (
	"log"
	"net/http"
	"time"
)

func Middlewarelog(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			next.ServeHTTP(w, r)
			log.Printf("method: %s, path: %s, duration: %s", r.Method, r.URL.Path, time.Since(startTime))
		},
	)
}
