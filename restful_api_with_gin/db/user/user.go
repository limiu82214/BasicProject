package user

import (
	"log"
	"strconv"

	"github.com/limiu82214/GoBasicProject/restful_api_with_gin/myutil/db"
	"github.com/limiu82214/GoBasicProject/restful_api_with_gin/myutil/gob"
)

type user struct {
	Name string `json:"name"`
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func NewUser() *user {
	u := &user{}
	return u
}

func GetUser(uid int) (u *user, err error) {
	u = NewUser()

	db := db.GetInst()
	data, _ := db.Get([]byte(`user/`+strconv.Itoa(uid)), nil)
	if data == nil {
		return nil, nil
	} else {

		err = gob.GetStrutFromByte(data, &u)
		if err != nil {
			log.Println(err)
			return nil, nil
		}
		return u, nil
	}
}

func CreateUser(uid int, u *user) (err error) {
	bytes, err := gob.StoreStructToByte(u)
	if err != nil {
		log.Fatalln(err)
	}

	db := db.GetInst()
	err = db.Put([]byte(`user/`+strconv.Itoa(uid)), bytes, nil)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func DeleteUser(uid int) (err error) {
	db := db.GetInst()
	err = db.Delete([]byte(`user/`+strconv.Itoa(uid)), nil)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
