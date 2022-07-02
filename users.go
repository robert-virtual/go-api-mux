package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func getUsers(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	json := json.NewEncoder(res)
	json.Encode(users)
}

func postUser(res http.ResponseWriter, req *http.Request) {
	js := json.NewEncoder(res)
	var user User
	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil {
		js.Encode(map[string]string{"err": err.Error()})
		return
	}
	users = append(users, user)
	res.WriteHeader(http.StatusOK)
	js.Encode(users)
}
