package db

type user struct {
	Name string `json:"name"`
}

func NewUser() *user {
	u := &user{}
	return u
}
