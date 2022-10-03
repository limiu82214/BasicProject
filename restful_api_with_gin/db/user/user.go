package user

import (
	"log"

	"github.com/limiu82214/GoBasicProject/restful_api_with_gin/myutil/db"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func GetUser(uid int) (u *User, err error) {
	u = &User{}
	d := db.GetInst()
	d.Take(&u, uid)
	if d.RowsAffected > 0 {
		return u, nil
	} else {
		return nil, nil
	}
}

func CreateUser(uid int, u *User) (err error) {
	db := db.GetInst()
	db.Create(u)
	return nil
}

func DeleteUser(uid int) (err error) {
	db := db.GetInst()
	db.Delete(uid)
	return nil
}
