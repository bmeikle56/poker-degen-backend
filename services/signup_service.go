package services

import (
	"pokerdegen/database"
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"regexp"
)

func SignupService(username string, password string) error {
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
	err = validatePassword(password)
	if err != nil {
		return err
	}

	err = database.InsertUser(db, username, string(hashed), 100)
	return err
}

func validatePassword(password string) error {
	// password must be 8+ characters
	if len(password) < 8 {
		return fmt.Errorf("password must be 8+ characters")
	}

	// password must contain a letter
	hasLetter := regexp.MustCompile(`[A-Za-z]`).MatchString
	
	// password must contain a number
	hasDigit := regexp.MustCompile(`\d`).MatchString

	if !hasLetter(password) {
		return fmt.Errorf("password must have a letter")
	} else if !hasDigit(password) {
		return fmt.Errorf("password must have a digit")
	}
	return nil
}