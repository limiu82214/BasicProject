package db

import (
	"bytes"
	"encoding/gob"
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
	db, err := leveldb.OpenFile("db/member", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	data, err := db.Get([]byte(`user/`+strconv.Itoa(uid)), nil)
	if data == nil {
		return nil, nil
	} else {
		enc := gob.NewDecoder(bytes.NewReader(data))
		err := enc.Decode(&u)
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

	var data bytes.Buffer
	enc := gob.NewEncoder(&data)
	enc.Encode(u)

	err = db.Put([]byte(`user/`+strconv.Itoa(uid)), data.Bytes(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
