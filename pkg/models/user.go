package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	IsVerified bool   `json:"is_verified" gorm:"default:false"`
	Name       string `json:"name"`
	Username   string `json:"Username" gorm:"type:varchar(255);unique"`
	Password   string `json:"Password" gorm:"type:varchar(255)"`
	Email      string `json:"Email" gorm:"type:varchar(255);unique"`
	Type       string `json:"Type" gorm:"type:enum('author','reader','admin')"`
	AuthorID   uint   `json:"author_id" gorm:"foreignkey:AuthorID"`
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func (b *User) CreateUser() *User {
	db.Create(&b)

	if b.Type == "author" {
		db.Last(&b)
		author := Author{Publisher: ""}
		db.Create(&author)
		b.AuthorID = author.ID
		db.Model(&author).Update("user_id", b.ID)
		db.Model(&b).Update("author_id", author.ID)
	}

	return b
}

func GetUserById(ID int64) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("ID=?", ID).Find(&getUser)
	return &getUser, db
}

func DeleteUser(ID int64) User {
	var author Author
	var user User
	db.Where("ID = ?", ID).Delete(&user)
	db.Where("ID = ?", author.UserID).Delete(&Author{})
	return user
}
