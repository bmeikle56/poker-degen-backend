package services

import (
	"pokerdegen/database"
	"golang.org/x/crypto/bcrypt"
)

func SignupService(username string, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	db, err := database.ConnectDB()
	if err != nil {
		return err
	}
	err = database.InsertUser(db, username, string(hashed), 100)
	return err
}