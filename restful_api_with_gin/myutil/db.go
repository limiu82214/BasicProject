package myutil

import (
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
		ShutdownServer(err)
	}
	return db
}
func DisconnectDB() {
	if db != nil {
		db.Close()
	}
}
