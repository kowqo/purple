package main

import (
	"fmt"
	"net/http"
	"purple/configs"
	"purple/internal/hello"
)

func main() {
	conf := configs.NewConfig()
	router := http.NewServeMux()
	hello.NewHelloHandler(router)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Listening on port 8080")
	server.ListenAndServe()
}
