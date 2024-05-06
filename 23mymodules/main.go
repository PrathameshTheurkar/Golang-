package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Modules in Golang!!")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", r))
}

func greeter() {
	fmt.Println("Hey there from golang!!")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hi from serveHome</h1>"))
}

// go mod tidy
// go mod verify
// go list
// go list all
// go list -m all
// go list -m -versions github.com/gorilla/mux
// go mod why github.com/gorilla/mux
// go mod graph
// go mod edit -go 1.16
// go mod edit -module "module_name"
// go mod vendor //can upgrade the module as per ur requirement
// go run -mod=vendor main.go //to use the upgraded version of module u need to use this command
