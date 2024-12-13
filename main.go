package main

import (
    "log"
    "net/http"
	"github.com/arcesti/PrimeiraAPIgo/user"
    "github.com/gorilla/mux"
)

func main() {
	// Inicializa o roteador
	r := mux.NewRouter()

	mockUsers = append(mockUsers, user.User{ID: "1", Nome: "Alice", Email: "alice@example.com"})
    mockUsers = append(mockUsers, user.User{ID: "2", Nome: "Bob", Email: "bob@example.com"})

	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/user/{id}", getUser).Methods("GET")
	r.HandleFunc("/users/create", createUser).Methods("POST")
	r.HandleFunc("/users/delete/{id}", deleteUser).Methods("DELETE")
	r.HandleFunc("/users/update/{id}", updateUser).Methods("PUT")

	log.Println("Servidor rodanod em http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}