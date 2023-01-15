package models

import (
	"errors"

	"gin-vue/utils/token"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID               int    `gorm:"primary_key" json:"id"`
	Username         string `gorm:"size:255;not null;unique" json:"username"`
	Email            string `gorm:"size:255;not null;unique" json:"email"`
	Password         string `gorm:"size:255;not null" json:"password"`
	Roleid           int    `json:"roleId"`
	Status           bool   `json:"status"`
	IsVerified       bool   `json:"is_verified"`
	Confirmationcode string `gorm:"not null" json:"confirmationCode"`
}

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (User) TableName() string {
	return "users"
}

func GetRoleName(rid int) string {

	var r Role
	if err := DB.First(&r, rid).Error; err != nil {
		return "Role not found!"
	}
	return r.Name
}

func GetUserByID(uid uint) (User, error) {

	var u User
	if err := DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}
	u.PrepareGive()
	return u, nil
}

func GetUserByEmail(email string) User {

	u := User{}
	DB.Model(User{}).Where("email = ?", email).Take(&u)
	return u
}

func (u *User) PrepareGive() {
	u.Password = ""
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(email string, password string) (string, error) {

	var err error
	u := User{}
	err = DB.Model(User{}).Where("email = ?", email).Take(&u).Error
	if err != nil {
		return "", err
	}
	err = VerifyPassword(password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	token, err := token.GenerateToken(u.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *User) SaveUser() (*User, error) {
	err := DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave() error {
	// turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
