package link

import (
	"fmt"
	"go-demo-6/pkg/request"
	"go-demo-6/pkg/response"
	"net/http"
)

type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
}

type HandlerLink struct {
	LinkRepository *LinkRepository
}

func NewHandlerLink(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &HandlerLink{
		LinkRepository: deps.LinkRepository,
	}

	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("GET /{hash}", handler.GoTo())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
}

func (handler *HandlerLink) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("create")

		body, err := request.HandleBody[LinkCreateRequest](&w, r)
		if err != nil {
			return
		}

		link := NewLink(body.Url)
		createdLink, err := handler.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.Json(w, createdLink, 200)
	}
}

func (handler *HandlerLink) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("read")

		hash := r.PathValue("hash")

		fmt.Println(hash)

		link, err := handler.LinkRepository.GetByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println(link.Url)

		// response.Json(w, "Read successful", 200)
		w.Header().Set("Content-Type", "application/json")
		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}

func (handler *HandlerLink) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("update")

		id := r.PathValue("id")
		fmt.Println(id)

		response.Json(w, "Update successful", 200)
	}
}

func (handler *HandlerLink) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("delete")

		id := r.PathValue("id")
		fmt.Println(id)

		response.Json(w, "Delete successful", 200)
	}
}
