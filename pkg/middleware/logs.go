package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapper := &WrapperWrite{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		start := time.Now()
		next.ServeHTTP(wrapper, r)
		log.Println(wrapper.StatusCode, r.Method, r.URL.Path, time.Since(start))
	})
}
