// main.go (API Todo List)
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type Todo struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

var (
	todos []Todo
	mu    sync.Mutex
	idSeq int
)

// Middleware CORS
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			return
		}
		next(w, r)
	}
}

// Handler
func getTodos(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	json.NewEncoder(w).Encode(todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	idSeq++
	todo.ID = idSeq
	todos = append(todos, todo)
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	var updatedTodo Todo
	json.NewDecoder(r.Body).Decode(&updatedTodo)

	mu.Lock()
	defer mu.Unlock()

	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Task = updatedTodo.Task
			todos[i].Completed = updatedTodo.Completed
			json.NewEncoder(w).Encode(todos[i])
			return
		}
	}
	http.NotFound(w, r)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	mu.Lock()
	defer mu.Unlock()

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}

func main() {
	// Tambah data awal
	todos = []Todo{
		{ID: 1, Task: "Belajar Golang", Completed: false},
	}
	idSeq = 1

	http.HandleFunc("/todos", enableCORS(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getTodos(w, r)
		case "POST":
			createTodo(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))

	http.HandleFunc("/todos/update", enableCORS(updateTodo))
	http.HandleFunc("/todos/delete", enableCORS(deleteTodo))

	fmt.Println("API Todo berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
