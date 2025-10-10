package main

import (
	"go-game-project/handlers"
	"log"
	"net/http"
)

func main() {

	router := http.NewServeMux()

	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	port := ":8080"
	log.Printf("Server running...")
	log.Fatal(http.ListenAndServe(port, router))
}
