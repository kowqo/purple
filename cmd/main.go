package main

import (
	"fmt"
	"net/http"
	"purple/configs"
	"purple/internal/auth"
	"purple/internal/links"
	"purple/pkg/db"
	"purple/pkg/middleware"
)

func main() {
	conf := configs.NewConfig()
	router := http.NewServeMux()
	db := db.NewDb(conf)

	//Repository
	linkRepository := links.NewLinkRepository(db)

	//handler
	links.NewLinkHandler(router, links.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	// Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}

	fmt.Println("Listening on port 8080")
	server.ListenAndServe()
}
