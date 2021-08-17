package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Status struct {
	Status string
}

func GetHomePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "UP")
}

func GetStatus(w http.ResponseWriter, r *http.Request){
	var status = Status{
		"OK",
	}
	err := json.NewEncoder(w).Encode(status)
	if err != nil {
		return
	}
}


func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", GetHomePage)
	myRouter.HandleFunc("/status", GetStatus)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()
}
