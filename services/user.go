package services

import (
	"golang.org/x/crypto/bcrypt"

	"e-commerce/utilities"
	"e-commerce/objects"
	"e-commerce/models"
)

var db = utilities.ConnecDB()

func GetUserByEmail(email string) (objects.User, error) {
	var user objects.User
	err := db.Where("email = ?", email).First(&user).Error
	return user, err
}


func CreateUser(user objects.Register) (objects.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return objects.User{}, err
	}

	newUser := models.User{
		Username: user.Username,
		Email:    user.Email,
		Password: string(hashedPassword),
	}

	if err := db.Create(&newUser).Error; err != nil {
		return objects.User{}, err
	}

	return objects.User{
		ID:       int(newUser.ID),
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func GetUserByID(id int) (objects.User, error) {
	var user objects.User
	err := db.Where("id = ?", id).First(&user).Error
	return user, err
}


func UpdateUserToken(id int, token string) error {
	return db.Model(&models.User{}).Where("id = ?", id).Update("token", token).Error
}