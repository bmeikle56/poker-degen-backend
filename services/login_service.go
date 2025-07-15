package services

import (
	"pokerdegen/database"
	"golang.org/x/crypto/bcrypt"
)

func LoginService(username string, password string) error {
	db, err := database.ConnectDB()
	if err != nil {
		return err
	}
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	_, err = database.FetchUser(db, username, string(hashed))
	return err
}