package services

import (
	"pokerdegen/database"
	"fmt"
	"pokerdegen/utils"
)

func DeleteAccountService(username string, password string) error {
	// connect to db
	db := database.GetDB()

	// validate username + password
	hashedPassword, err := database.FetchPasswordForUser(db, username)
	if err != nil {
		return err
	}
	err = utils.ComparePassword(hashedPassword, password)
	if err != nil {
		return fmt.Errorf("incorrect password")
	}

	// now we know username + password are valid, delete account
	err = database.DeleteUser(db, username)
	return err
}