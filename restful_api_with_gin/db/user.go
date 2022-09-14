package db

type user struct {
	Name string `json:"name"`
}

func NewUser() *user {
	u := &user{}
	return u
}

func GetUser(uid int) (u *user, err error) {
	u = NewUser()
	switch uid {
	case 1:
		u.Name = "mike" // 應在資料庫存取 *todo*
	case 2:
		u.Name = "joe"
	}
	return u, nil
}
