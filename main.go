package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/gorilla/mux"
)

type Status struct {
	Status string
}

type Test struct {
	description string
}

var ITEMS = 10000
var testArray [10000]Test

func GetHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UP")
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
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

	myRouter.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)

	// register router manually
	//myRouter.HandleFunc("/debug/pprof/", pprof.Index)
	//myRouter.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	//myRouter.HandleFunc("/debug/pprof/profile", pprof.Profile)
	//myRouter.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	//myRouter.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	//myRouter.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	//myRouter.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	//myRouter.Handle("/debug/pprof/block", pprof.Handler("block"))

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func populateArray() {
	for i := 0; i < ITEMS; i++ {
		testArray[i] = Test{description: "Test " + string(i)}
	}
}

func main() {
	populateArray()
	handleRequests()
}
