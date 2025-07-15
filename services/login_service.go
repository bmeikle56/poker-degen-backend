package services

import (
	"pokerdegen/database"
)

func LoginService(username string, password string) bool {
	db, _ := database.ConnectDB()
	_, err := database.FetchUser(db, username, password)
	return err == nil
}