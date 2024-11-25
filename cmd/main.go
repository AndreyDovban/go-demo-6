package main

import (
	"fmt"
	"go-demo-6/configs"
	"go-demo-6/internal/auth"
	"go-demo-6/internal/link"
	"go-demo-6/internal/stat"
	"go-demo-6/internal/user"
	"go-demo-6/pkg/db"
	"go-demo-6/pkg/event"
	"go-demo-6/pkg/middleware"
	"net/http"
)

func main() {

	config := configs.LoadConfig()
	db := db.NewDb(config)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	// Repositoryes
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepository := stat.NewStatrepository(db)

	// Services
	authService := auth.NewAuthService(userRepository)
	statService := stat.NewStatService(&stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepository,
	})

	// Handlers
	auth.NewHandlerAuth(router, auth.AuthHandlerDeps{
		Config:      config,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		Config:         config,
		EventBus:       eventBus,
	})
	stat.NewStatHandler(router, &stat.StatHandlerDeps{
		StatRepository: statRepository,
		Config:         config,
	})

	// Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := &http.Server{
		Addr:    ":3000",
		Handler: stack(router),
	}

	go statService.AddClick()

	fmt.Println("http://localhost:3000")
	server.ListenAndServe()
}

// Context with cancel
// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	go tickOperation(ctx)
// 	time.Sleep(2 * time.Second)
// 	cancel()
// }
// func tickOperation(ctx context.Context) {
// 	ticker := time.NewTicker(200 * time.Millisecond)
// 	for {
// 		select {
// 		case <-ticker.C:
// 			fmt.Println("tick")
// 		case <-ctx.Done():
// 			fmt.Println("Cancel")
// 			return
// 		}
// 	}
// }

// Context with timeoute
// func main() {
// 	ctx := context.Background()
// 	ctxWithTimout, cancel := context.WithTimeout(ctx, 4*time.Second)
// 	defer cancel()
// 	done := make(chan struct{})
// 	go func() {
// 		time.Sleep(3 * time.Second)
// 		close(done)
// 	}()
// 	select {
// 	case <-done:
// 		fmt.Println("Done task")
// 	case <-ctxWithTimout.Done():
// 		fmt.Println("timoute")
// 	}
// }

// Context with key-value
// func main() {
// 	type key int
// 	const EmailKey key = 0
// 	ctx := context.Background()
// 	ctxWithValue := context.WithValue(ctx, EmailKey, "a@a.ru")
// 	if userEmail, ok := ctxWithValue.Value(EmailKey).(string); ok {
// 		fmt.Println(userEmail)
// 	} else {
// 		fmt.Println("No value")
// 	}
// }
