package main

import (
	"fmt"
	"net/http"

	"github.com/nafnaufal/roadmap-forum/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/favicon.ico", handlers.FaviconHandler)

	fmt.Println("Server started at: http://localhost:8080")
	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println(err.Error())
	}
}
