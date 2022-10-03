package user

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/limiu82214/GoBasicProject/restful_api_with_gin/myutil/db"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func GetUser(uid int) (u *User, err error) {
	u = &User{}
	d := db.GetInst()
	rst := d.Take(&u, uid)
	if rst.RowsAffected > 0 {
		return u, nil
	} else {
		return nil, nil
	}
}

func CreateUser(u *User) (uid int, err error) {
	db := db.GetInst()
	db.Create(&u)
	return u.Uid, nil
}

func DeleteUser(uid int) (err error) {
	db := db.GetInst()
	u := &User{
		Uid: uid,
	}
	db.Delete(u)
	return nil
}
