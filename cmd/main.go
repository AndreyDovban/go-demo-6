package main

import (
	"fmt"
	"go-demo-6/configs"
	"go-demo-6/internal/auth"
	"go-demo-6/internal/link"
	"go-demo-6/internal/user"
	"go-demo-6/pkg/db"
	"go-demo-6/pkg/middleware"
	"net/http"
)

func main() {

	config := configs.LoadConfig()
	db := db.NewDb(config)
	router := http.NewServeMux()

	// Repositoryes
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)

	// Services
	authService := auth.NewAuthService(userRepository)

	//
	auth.NewHandlerAuth(router, auth.AuthHandlerDeps{Config: config, AuthService: authService})
	link.NewHandlerLink(router, link.LinkHandlerDeps{LinkRepository: linkRepository})

	// Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := &http.Server{
		Addr:    ":3000",
		Handler: stack(router),
	}

	fmt.Println("http://localhost:3000")
	server.ListenAndServe()
}
