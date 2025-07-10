package services

func LoginService(username string, password string) bool {
	// for now, just always return true
	if (username == "braeden" && password == "pokerdegen") {
		return true
	}
	return false
}