package main

import (
	"fmt"
	"net/http"

	"go-demo-6/configs"

	auth "go-demo-6/internal/auth"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	fmt.Println("http://localhost:3000")
	server.ListenAndServe()
}
