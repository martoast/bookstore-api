package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"bookstore-api/pkg/models"

	"bookstore-api/pkg/utils"

	"github.com/gorilla/mux"
)

var NewAuthor models.Author

func GetAuthor(w http.ResponseWriter, r *http.Request) {
	newAuthors := models.GetAllAuthors()

	res, err := json.Marshal(newAuthors)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	CreateAuthor := &models.Author{}
	utils.ParseBody(r, CreateAuthor)
	b := CreateAuthor.CreateAuthor()
	res, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAuthorById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	AuthorId := vars["AuthorId"]
	ID, err := strconv.ParseInt(AuthorId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	AuthorDetails, _ := models.GetAuthorById(ID)
	res, _ := json.Marshal(AuthorDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	AuthorId := vars["AuthorId"]
	ID, err := strconv.ParseInt(AuthorId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	Author := models.DeleteAuthor(ID)
	res, _ := json.Marshal(Author)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	var updateAuthor = &models.Author{}
	utils.ParseBody(r, updateAuthor)
	vars := mux.Vars(r)
	AuthorId := vars["AuthorId"]
	ID, err := strconv.ParseInt(AuthorId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	AuthorDetails, db := models.GetAuthorById(ID)
	if updateAuthor.Publisher != "" {
		AuthorDetails.Publisher = updateAuthor.Publisher
	}

	db.Save(&AuthorDetails)
	res, err := json.Marshal(AuthorDetails)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
