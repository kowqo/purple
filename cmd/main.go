package main

import (
	"fmt"
	"net/http"
	"purple/internal/auth"
)

func main() {
	//conf := configs.NewConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(router)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Listening on port 8080")
	server.ListenAndServe()
}
