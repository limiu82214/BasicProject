package db

import (
	"github.com/jinzhu/gorm"
	"github.com/limiu82214/GoBasicProject/restful_api_with_gin/myutil/sig"
)

var db *gorm.DB
var err error

func init() {
}

func GetInst() *gorm.DB {
	if db != nil {
		return db
	}

	dsn := "root:root@tcp(127.0.0.1:3306)/dbtest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		sig.ShutdownServer(err)
	}
	db.LogMode(true)
	return db
}
func DisconnectDB() {
	if db != nil {
		db.Close()
	}
}
