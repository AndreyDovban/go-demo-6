package main

import (
	"fmt"
	"go-demo-6/configs"
	"go-demo-6/internal/auth"
	"go-demo-6/pkg/db"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	_ = db.NewDb(config)
	router := http.NewServeMux()
	auth.NewHandlerAuth(router, config)

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	fmt.Println("http://localhost:3000")
	server.ListenAndServe()
}
