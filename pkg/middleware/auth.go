package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")
		token := strings.TrimPrefix(bearer, "Bearer")
		fmt.Println(token)
		if len(token) == 0 {
			http.Error(w, errors.New("No Token Provided").Error(), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
