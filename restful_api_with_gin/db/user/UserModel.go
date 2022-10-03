package user

type User struct {
	Uid     uint `gorm:"primary_key"`
	Account string
	Pwd     string
	Name    string
	Age     int
}
