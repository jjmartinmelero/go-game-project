package handlers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home Page")
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Crear nuevo juego")
}

func Game(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "juego")
}

func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "jugar")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "about")
}
