package auth

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"purple/configs"
	"purple/pkg/resp"
	"regexp"
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
		var payload LoginRequest
		err := json.NewDecoder(req.Body).Decode(&payload)

		if err != nil {
			resp.JSON(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = validateLogin(payload)
		if err != nil {
			resp.JSON(w, err.Error(), http.StatusBadRequest)
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
		w.Write([]byte("register"))
		log.Println(h.Auth.Secret)
	}
}

func validateLogin(l LoginRequest) error {
	if len(l.Email) == 0 || len(l.Password) == 0 {
		return errors.New("email or password is empty")
	}
	reg, _ := regexp.Compile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	if !reg.MatchString(l.Email) {
		return errors.New("invalid email")
		return
	}
	return nil
}
