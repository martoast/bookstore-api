package models

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"Username" gorm:"type:varchar(255);unique"`
	Password string `json:"Password" gorm:"type:varchar(255)"`
	Email    string `json:"Email" gorm:"type:varchar(255);unique"`
	Type     string `json:"Type" gorm:"type:enum('author','reader','admin')"`
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}
