package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

func main() {

	config := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "godb",
		AllowNativePasswords:true,
	}
	var err error
	db, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

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
