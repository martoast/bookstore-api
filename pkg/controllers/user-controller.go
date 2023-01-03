package controllers

import (
	"encoding/json"

	"net/http"

	"bookstore-api/pkg/models"
)

var NewUser models.User

func GetUser(w http.ResponseWriter, r *http.Request) {
	newUsers := models.GetAllUsers()

	res, err := json.Marshal(newUsers)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
