package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)
var users []User
func main() {
	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router))
	
	usersRouter := router.PathPrefix("/users").Subrouter()
	usersRouter.HandleFunc("", postUser).Methods("POST")
	usersRouter.HandleFunc("", getUsers).Methods("GET")
	
	http.Handle("/", router)
	server := &http.Server{
		Handler:      router,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())	
}
