package auth

import (
	"fmt"
	"go-demo-6/configs"
	"go-demo-6/pkg/jwt"
	"go-demo-6/pkg/request"
	"go-demo-6/pkg/response"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

type HandlerAuth struct {
	*configs.Config
	*AuthService
}

func NewHandlerAuth(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &HandlerAuth{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *HandlerAuth) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}

		email, err := handler.AuthService.Login(body.Email, body.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		fmt.Println(handler.Config.Auth.Secret)
		j := jwt.NewJWT(handler.Config.Auth.Secret)
		s, err := j.Create(jwt.JWTData{
			Email: email,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data := LoginResponse{
			Token: s,
		}

		response.Json(w, data, 200)
	}
}

func (handler *HandlerAuth) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}

		email, err := handler.AuthService.Register(body.Email, body.Password, body.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		j := jwt.NewJWT(handler.Config.Auth.Secret)
		s, err := j.Create(jwt.JWTData{
			Email: email,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data := LoginResponse{
			Token: s,
		}

		response.Json(w, data, 200)
	}
}
