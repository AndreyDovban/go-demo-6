package auth

import (
	"fmt"
	"go-demo-6/configs"
	"net/http"

	resp "go-demo-6/pkg/response"
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

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		secret := handler.Config.Auth.Secret
		fmt.Println("Login ", secret)
		data := LoginResponse{
			Token: "123 !!!",
		}

		resp.Json(w, data, 202)

	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Register")

	}
}
