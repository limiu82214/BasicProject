package user

type User struct {
	Uid     int `gorm:"primary_key"`
	Account string
	Pwd     string
	Name    string
	Age     int
}
