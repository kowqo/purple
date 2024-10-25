package auth

import (
	"fmt"
	"net/http"
)

type AuthHandler struct {
}

func NewAuthHandler(router *http.ServeMux) {
	handler := &AuthHandler{}

	router.HandleFunc("/login", handler.Login())
	router.HandleFunc("/register", handler.Register)
}

func (h *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello Login")
	}
}
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {}
