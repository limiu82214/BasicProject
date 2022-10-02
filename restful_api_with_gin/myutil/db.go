package myutil

import (
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB
var err error

func GetInst() *leveldb.DB {
	if db != nil {
		return db
	}
	db, err = leveldb.OpenFile("db/member", nil)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
