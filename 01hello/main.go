package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var users []User

func main() {
	fmt.Println("Hello from the first server")

	// SEEDING
	users = append(users, User{"1", "Prathamesh"})
	users = append(users, User{"2", "Prajwal"})
	users = append(users, User{"3", "Prajakta"})
	users = append(users, User{"4", "Shubham"})

	// ROUTER
	r := mux.NewRouter()

	// ROUTES
	r.HandleFunc("/", homePage).Methods("GET")
	r.HandleFunc("/users", getAllUsers).Methods("GET")
	r.HandleFunc("/user", createUser).Methods("POST")
	r.HandleFunc("/user/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")

	fmt.Println("Server is running on port 3000")
	// SERVER RUNNING
	http.ListenAndServe(":3000", r)

}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<h1>This is the home page . Welcome to home page</h1>`))
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var updatedUser User

	_ = json.NewDecoder(r.Body).Decode(&updatedUser)

	params := mux.Vars(r)

	for index, user := range users {
		if user.Id == params["id"] {
			users = append(users[:index], users[index+1:]...)
			users = append(users, updatedUser)
			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}

	json.NewEncoder(w).Encode("No user found!!")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, user := range users {
		if user.Id == params["id"] {
			users = append(users[:index], users[index+1:]...)
			json.NewEncoder(w).Encode("User deleted successfully")
			return
		}
	}

	json.NewEncoder(w).Encode("No user found")
}
