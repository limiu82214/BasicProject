package db

import (
	"log"
	"strconv"

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
	db, err := leveldb.OpenFile("./member", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	data, err := db.Get([]byte(`user/`+strconv.Itoa(uid)), nil)
	if data == nil {
		return nil, nil
	} else {
		u.Name = string(data)
		return u, nil
	}
}

func CreateUser(uid int, u *user) (err error) {
	db, err := leveldb.OpenFile("./member", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Put([]byte(`user/`+strconv.Itoa(uid)), []byte(u.Name), nil)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
