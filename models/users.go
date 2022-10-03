package models

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	"golang-mysql-restful-starter-kit/config"
	"golang-mysql-restful-starter-kit/helpers"
)

type User struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAT time.Time `json:"createdAt"`
}

func NewUser(name, email, password string) (*User, error) {

	u := &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	u.CreatedAT = time.Now()
	fmt.Println(u)
	// validating name field with retuired, min length 3, max length 25 and no regex check
	if e := helpers.Validator(u.Name, true, 3, 25, "", "Name"); e != nil {
		return nil, e
	}

	// validating email field with required, min length 5, max length 25 and regex check
	if e := helpers.Validator(u.Email, true, 5, 25, `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, "Email"); e != nil {
		return nil, e
	}

	// validating password field with required, min length 8, max length 25 and regex check
	if e := helpers.Validator(u.Password, true, 8, 25, "^[a-zA-Z0-9_!@#$_%^&*.?()-=+]*$", "Password"); e != nil {
		return nil, e
	}

	passwordBytes := []byte(password + os.Getenv(config.SALT))
	hash, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	u.Password = string(hash[:])
	if err != nil {
		return nil, err
	}

	return u, nil
}

// ComparePassword , used to compared
// hashed password with input text password
// return error if any otherwise nil
func (u *User) ComparePassword(password string) error {
	incoming := []byte(password + os.Getenv(config.SALT))
	existing := []byte(u.Password)
	err := bcrypt.CompareHashAndPassword(existing, incoming)
	return err
}
