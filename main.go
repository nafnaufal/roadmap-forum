package main

import (
	"fmt"
	"net/http"

	"github.com/nafnaufal/roadmap-forum/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.RootHandler)
	mux.HandleFunc("GET /favicon.ico", handlers.FaviconHandler)

	fmt.Println("Server started at: http://localhost:8080")
	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println(err.Error())
	}
}
