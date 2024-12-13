package main

import (
	"encoding/json"
	"net/http"
	"github.com/arcesti/PrimeiraAPIgo/db"
	"github.com/arcesti/PrimeiraAPIgo/user"
	"github.com/gorilla/mux"
)

var mockUsers []user.User


func getUsers(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type","Application/json")
	json.NewEncoder(res).Encode(mockUsers)
}

func getUser(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for _,user := range mockUsers {
		if user.ID == params["id"] {
			res.Header().Set("Content-Type", "Application/json")
			json.NewEncoder(res).Encode(user)
			return
		}
	}
	http.Error(res, "User not found", http.StatusNotFound)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user user.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err!=nil {
		http.Error(w,"Erro ao criar usuario", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
	err = db.Create(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for usrId, usrDel := range mockUsers{
		if usrDel.ID == params["id"] {
			mockUsers = append(mockUsers[:usrId], mockUsers[usrId+1:]...)
			w.Header().Set("Content-Type","Application/json")
			json.NewEncoder(w).Encode(mockUsers)
			return
		}
	}
	http.Error(w, "Erro ao deletar usuario", http.StatusBadRequest)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var newUsr user.User
	err := json.NewDecoder(r.Body).Decode(&newUsr)
	if(err!=nil) {
		http.Error(w, "Usuario invalido", http.StatusBadRequest)
		return
	}
	for index,user := range mockUsers {
		if user.ID == params["id"] {
			mockUsers[index] = newUsr
			w.Header().Set("Content-Type", "Application/json")
			json.NewEncoder(w).Encode(mockUsers)
			return
		}
	}
	http.Error(w, "Erro ao alterar usuario", http.StatusBadRequest)
}