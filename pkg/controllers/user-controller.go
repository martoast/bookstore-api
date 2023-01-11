package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"net/http"

	"bookstore-api/pkg/models"
	"bookstore-api/pkg/utils"

	"github.com/gorilla/mux"
)

var NewUser models.User

type CreateUserRequest struct {
	User   *models.User   `json:"user"`
	Author *models.Author `json:"author,omitempty"`
}

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

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	b := CreateUser.CreateUser()
	res, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UserId := vars["UserId"]
	ID, err := strconv.ParseInt(UserId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	UserDetails, _ := models.GetUserById(ID)
	res, _ := json.Marshal(UserDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UserId := vars["UserId"]
	ID, err := strconv.ParseInt(UserId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	User := models.DeleteUser(ID)
	res, _ := json.Marshal(User)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser = &models.User{}
	utils.ParseBody(r, updateUser)
	vars := mux.Vars(r)
	UserId := vars["UserId"]
	ID, err := strconv.ParseInt(UserId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	UserDetails, db := models.GetUserById(ID)
	if updateUser.Email != "" {
		UserDetails.Email = updateUser.Email
	}
	if updateUser.Password != "" {
		UserDetails.Password = updateUser.Password
	}

	db.Save(&UserDetails)
	res, err := json.Marshal(UserDetails)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
