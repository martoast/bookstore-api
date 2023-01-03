package models

import (
	"bookstore-api/pkg/config"
	"fmt"

	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	AuthorID    uint   `json:"author_id" gorm:"foreignkey:AuthorID"`
	Author      Author
	Publication string `json:"publication"`
	Cover       string `json:"cover"`
}

func init() {
	config.Connect()
	db = config.GetDB()

	db.AutoMigrate(&Book{})
	db.AutoMigrate(&Author{})
	db.AutoMigrate(&User{})
}

func GetAllBooks() []Book {
	var Books []Book
	db.Preload("Author").Find(&Books)
	return Books
}

func (b *Book) CreateBook() *Book {
	var author Author
	if err := db.Where("ID = ?", b.AuthorID).First(&author).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// The referenced author record was not found
			fmt.Println("Referenced author record was not found.")
		} else {
			fmt.Println(err)
		}
	} else {
		// The referenced author record was found
		db.Model(&b).Association("Author").Append(b.Author)
	}
	db.Create(&b)
	return b
}

func GetBookById(ID int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Preload("Author").Where("ID=?", ID).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID = ?", ID).Delete(&book)
	db.Where("ID = ?", book.AuthorID).Delete(&Author{})
	return book
}
