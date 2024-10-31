package link

import (
	"fmt"
	"go-demo-6/configs"
	"go-demo-6/pkg/response"
	"net/http"
)

type LinkHandlerDeps struct {
	Config *configs.Config
}

type HandlerLink struct {
	Config *configs.Config
}

func NewHandlerLink(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &HandlerLink{
		Config: deps.Config,
	}

	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("GET /link/{alias}", handler.GoTo())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
}

func (handler *HandlerLink) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("create")

		response.Json(w, "Create successful", 200)
	}
}

func (handler *HandlerLink) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("read")

		response.Json(w, "Read successful", 200)
	}
}

func (handler *HandlerLink) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("update")

		response.Json(w, "Update successful", 200)
	}
}

func (handler *HandlerLink) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("delete")

		response.Json(w, "Delete successful", 200)
	}
}
