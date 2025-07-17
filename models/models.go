package models

type Board struct {
	CC1 string `json:"cc1"`
	CC2 string `json:"cc2"`
	CC3 string `json:"cc3"`
	CC4 string `json:"cc4"`
	CC5 string `json:"cc5"`

	V1C1 string `json:"v1c1"`
	V1C2 string `json:"v1c2"`

	HC1 string `json:"hc1"`
	HC2 string `json:"hc2"`

	POT string `json:"pot"`
}

type ModelRequest struct {
	Board Board `json:"board"`
}

type User struct {
	ID       string
	Username string
	Password string
	Diamonds int
}

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
