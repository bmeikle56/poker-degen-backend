package services

import (
	"pokerdegen/database"
	"fmt"
	"pokerdegen/utils"
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
	err = utils.ComparePassword(hashedPassword, password)
	if err != nil {
		return fmt.Errorf("incorrect password")
	}
	_, err = database.FetchUser(db, username)
	return err
}