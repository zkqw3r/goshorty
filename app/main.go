package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
)

var cfg = LoadConfig()

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env not found")
	}

	cfg := LoadConfig()

	InitDB(cfg.DatabaseURL)
	defer db.Close()
	fmt.Println("Successfully connected to the database!")

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.HandleFunc("/", handleMain)
	http.HandleFunc("/shorten", handleShorten)
	http.HandleFunc("/stats/", handleStats)

	fmt.Printf("Server started: http://localhost%s\n", cfg.Port)
	err := http.ListenAndServe(cfg.Port, nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
