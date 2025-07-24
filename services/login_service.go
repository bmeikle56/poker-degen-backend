package services

import (
	"pokerdegen/database"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

func LoginService(username string, password string) error {
	db, err := database.ConnectDB()
	if err != nil {
		return err
	}
	hashedPassword, err := database.FetchPasswordForUser(db, username)
	if err != nil {
		return err
	}
	err = comparePassword(hashedPassword, password)
	if err != nil {
		return fmt.Errorf("incorrect password")
	}
	_, err = database.FetchUser(db, username)
	return err
}

func comparePassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}