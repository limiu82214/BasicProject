package leveldb

import (
	"github.com/limiu82214/GoBasicProject/restful_api_with_gin/myutil/sig"
	level_db "github.com/syndtr/goleveldb/leveldb"
)

var db_once *level_db.DB //nolint // signletone use name with _
var err error

func GetInst() *level_db.DB {
	if db_once != nil {
		return db_once
	}

	db_once, err = level_db.OpenFile("./assets/leveldb", nil)
	if err != nil {
		sig.ShutdownServer(err)
	}

	return db_once
}
func DisconnectDB() {
	if db_once != nil {
		db_once.Close()
	}
}
