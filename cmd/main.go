package main

import (
	"fmt"
	"net/http"
	"purple/configs"
	"purple/internal/auth"
	"purple/internal/links"
	"purple/pkg/db"
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

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Listening on port 8080")
	server.ListenAndServe()
}
