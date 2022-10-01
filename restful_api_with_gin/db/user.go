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
	db.Put([]byte(`user/1`), []byte(`mike`), nil) // *todo* 在測試中用post新增
	db.Put([]byte(`user/2`), []byte(`joe`), nil)

	data, err := db.Get([]byte(`user/`+strconv.Itoa(uid)), nil)
	if data == nil {
		return nil, nil
	} else {
		u.Name = string(data)
		return u, nil
	}
}
