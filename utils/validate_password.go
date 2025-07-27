package utils

import (
	"fmt"
	"regexp"
)

func ValidatePassword(password string) error {
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