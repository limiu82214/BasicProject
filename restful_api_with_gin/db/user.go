package db

import (
	"log"
	"strconv"

	"github.com/limiu82214/GoBasicProject/restful_api_with_gin/myutil"
	"github.com/syndtr/goleveldb/leveldb"
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
	db, err := leveldb.OpenFile("db/member", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	data, err := db.Get([]byte(`user/`+strconv.Itoa(uid)), nil)
	if data == nil {
		return nil, nil
	} else {

		err := myutil.GetStrutFromByte(data, &u)
		if err != nil {
			log.Println(err.Error())
			return nil, nil
		}
		return u, nil
	}
}

func CreateUser(uid int, u *user) (err error) {
	db, err := leveldb.OpenFile("db/member", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	bytes, err := myutil.StoreStructToByte(u)
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Put([]byte(`user/`+strconv.Itoa(uid)), bytes, nil)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
