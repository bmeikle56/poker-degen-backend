package services

import (
	"pokerdegen/database"
	"golang.org/x/crypto/bcrypt"
)

func SignupService(username string, password string) bool {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	db, _ := database.ConnectDB()
	err := database.InsertUser(db, username, string(hashed), 100)
	return err == nil
}