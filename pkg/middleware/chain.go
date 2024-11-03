package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func Chain(Middlewares ...Middleware) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		for _, middleware := range Middlewares {
			next = middleware(next)
		}

		return next
	}
}
