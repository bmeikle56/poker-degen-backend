package services

import (
	"pokerdegen/database"
)

func LoginService(username string, password string) error {
	db, err := database.ConnectDB()
	if err != nil {
		return err
	}
	_, err = database.FetchUser(db, username, password)
	return err
}