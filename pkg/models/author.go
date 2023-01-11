package models

import (
	"gorm.io/gorm"
)

type Author struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Publisher string `json:"publisher"`
	UserID    uint   `json:"user_id" gorm:"foreignkey:UserID"`
}

func GetAllAuthors() []Author {
	var Authors []Author
	db.Find(&Authors)
	return Authors
}

func (b *Author) CreateAuthor() *Author {
	db.Create(&b)
	return b
}

func GetAuthorById(ID int64) (*Author, *gorm.DB) {
	var getAuthor Author
	db := db.Where("ID=?", ID).Find(&getAuthor)
	return &getAuthor, db
}

func DeleteAuthor(ID int64) Author {
	var author Author
	db.Delete(&Author{}, ID)
	return author
}
