package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/nafnaufal/roadmap-forum/internal/db"
	"github.com/nafnaufal/roadmap-forum/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db.InitDB()
	defer db.CloseDB()

	mux.HandleFunc("GET /", handlers.RootHandler)
	mux.HandleFunc("GET /favicon.ico", handlers.FaviconHandler)

	mux.HandleFunc("GET /users", handlers.GetUsersHandler)
	mux.HandleFunc("POST /register", handlers.RegisterHandler)

	fmt.Println("Server started at: http://localhost:8080")
	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println(err.Error())
	}
}
