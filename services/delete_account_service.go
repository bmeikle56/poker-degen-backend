package services

import (
	"pokerdegen/database"
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"pokerdegen/utils"
)

func DeleteAccountService(username string, password string) error {
	// hash the user's password
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// connect to the database
	db, err := database.ConnectDB()
	if err != nil {
		return err
	}

	// now make sure the user doesn't already exist
	userExists, err := database.CheckIfUserExists(db, username)
	if err != nil {
		return err
	} else if userExists {
		return fmt.Errorf("user already exists")
	}

	// now we need to make sure username and password are valid + strong
	err = utils.ValidatePassword(password)
	if err != nil {
		return err
	}

	err = database.InsertUser(db, username, string(hashed), 100)
	return err
}