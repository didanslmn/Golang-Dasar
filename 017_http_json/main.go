package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Nama  string `json:"nama"`
	Email string `json:"email"`
	Umur  int    `json:"usia"`
}

// Handler HTTP
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Selamat datang di API Golang!")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		Nama:  "Budi",
		Email: "budi@email.com",
		Umur:  20,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/user", userHandler)

	fmt.Println("Server berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
