package auth

import (
	"net/http"
	"purple/configs"
	"purple/pkg/request"
	"purple/pkg/resp"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}

	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (h *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		_, err := request.HandleBody[LoginRequest](&w, req)
		if err != nil {
			return
		}

		res := LoginResponse{
			Token: "12321",
		}

		resp.JSON(w, res, http.StatusOK)
	}
}
func (h *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		_, err := request.HandleBody[RegisterRequest](&w, req)
		if err != nil {
			return
		}
		res := LoginResponse{
			Token: "12321",
		}

		resp.JSON(w, res, http.StatusOK)
	}
}
