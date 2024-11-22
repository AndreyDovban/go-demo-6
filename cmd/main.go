package main

import (
	"fmt"
	"go-demo-6/configs"
	"go-demo-6/internal/auth"
	"go-demo-6/internal/link"
	"go-demo-6/pkg/db"
	"net/http"
)

func main() {

	config := configs.LoadConfig()
	db := db.NewDb(config)
	router := http.NewServeMux()
	linkRepository := link.NewLinkRepository(db)

	auth.NewHandlerAuth(router, auth.AuthHandlerDeps{Config: config})
	link.NewHandlerLink(router, link.LinkHandlerDeps{LinkRepository: linkRepository})

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	fmt.Println("http://localhost:3000")
	server.ListenAndServe()
}
