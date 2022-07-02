package main

import (
	"encoding/json"
	"net/http"
)

func getUsers(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	json := json.NewEncoder(res)
	user := User{}
	users, err := user.FindAll()
	if err != nil {
		json.Encode(map[string]string{"err": err.Error()})
		return
	}
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
	_, err = user.Create()
	if err != nil {
		js.Encode(map[string]string{"err": err.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
	js.Encode(user)
}
