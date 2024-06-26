package models

import (
	"gofiber-marketplace/src/configs"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=20"`
	Role     string `json:"role"`
	Verify   string `json:"verify"`
}

type UserVerification struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	Token  string `json:"token"`
}

type Register struct {
	Name     string `json:"name" validate:"required,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required,numeric,max=15"`
	Password string `json:"password" validate:"required,min=8,max=20"`
	Role     string `json:"role" validate:"oneof=seller customer"`
}

func SelectAllUsers() []*User {
	var users []*User
	configs.DB.Find(&users)
	return users
}

func SelectUserById(id int) *User {
	var user User
	configs.DB.First(&user, "id = ?", id)
	return &user
}

func SelectUserbyEmail(email string) *User {
	var user User
	configs.DB.First(&user, "email = ?", email)
	return &user
}

func SelectUserVerification(userId int, token string) *UserVerification {
	var userVerification UserVerification
	configs.DB.Where("user_id = ? AND token = ?", userId, token).First(&userVerification)
	return &userVerification
}

func CreateUser(user *User) (uint, error) {
	result := configs.DB.Create(&user)
	return user.ID, result.Error
}

func CreateUserVerification(userVerification *UserVerification) error {
	result := configs.DB.Create(&userVerification)
	return result.Error
}

func UpdateUser(id int, updatedUser *User) error {
	result := configs.DB.Model(&User{}).Where("id = ?", id).Updates(updatedUser)
	return result.Error
}

func UpdateUserbyEmail(email string, updatedUser *User) error {
	result := configs.DB.Model(&User{}).Where("email = ?", email).Updates(updatedUser)
	return result.Error
}

func UpdateUserVerify(id int) error {
	result := configs.DB.Model(&User{}).Where("id = ?", id).Update("verify", "true")
	return result.Error
}

func DeleteUser(id int) error {
	result := configs.DB.Delete(&User{}, "id = ?", id)
	return result.Error
}

func DeleteUsersVerification(id int, token string) error {
	result := configs.DB.Where("id = ? AND token = ?", id, token).Delete(&UserVerification{})
	return result.Error
}
