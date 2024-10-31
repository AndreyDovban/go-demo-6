package auth

import (
	"go-demo-6/configs"
	"go-demo-6/pkg/request"
	"go-demo-6/pkg/response"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type HandlerAuth struct {
	Config *configs.Config
}

func NewHandlerAuth(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &HandlerAuth{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())

}

func (handler *HandlerAuth) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request.HandleBody[LoginRequest](&w, r)

		response.Json(w, "Successful login", 200)
	}
}

func (handler *HandlerAuth) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request.HandleBody[RegisterRequest](&w, r)

		response.Json(w, "Successful register", 200)
	}
}
