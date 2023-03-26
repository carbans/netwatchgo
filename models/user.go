package models

import (
	"errors"

	"github.com/carbans/netdebug/utils/token"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email          string `json:"email" gorm:"unique;not null;"`
	Password       string `json:"password" gorm:"not null;"`
	Name           string `json:"name"`
	EmailConfirmed bool   `json:"emailConfirmed,omitempty"`
	Agents         []Agent
	ApiKeys        []ApiKey
}

func GetUserByID(uid uint) (User, error) {

	var user User

	if err := DB.First(&user, uid).Error; err != nil {
		return user, errors.New("User not found")
	}

	user.PrepareGive()

	return user, nil

}

func (u *User) PrepareGive() {
	u.Password = ""
}

// function to verify password
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(email string, password string) (string, error) {
	user := User{}

	var err = DB.Model(User{}).Where("email = ?", email).Take(&user).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(user.Password, password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (user *User) SaveUser() (*User, error) {

	var err = DB.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

// Function to add hashed password before to saving user into database, register callback to gorm
func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}
