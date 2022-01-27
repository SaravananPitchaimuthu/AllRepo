package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func handle(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing Request!")
	w.Write([]byte("OK"))
	log.Println("finished Processing Request")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handle)
	loggingHandler := handlers.LoggingHandler(os.Stdout, router)
	http.ListenAndServe(":8000", loggingHandler)
}
