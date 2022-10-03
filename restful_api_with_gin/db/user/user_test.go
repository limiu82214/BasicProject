package user

import (
	"testing"

	"github.com/limiu82214/GoBasicProject/restful_api_with_gin/myutil/db"
	"github.com/stretchr/testify/assert"
)

func init() {
	ResetDB()
}
func ResetDB() {
	d := db.GetInst()
	d.Exec("DROP TABLE users")
	d.Exec("CREATE TABLE `users` ( `uid` int(11) NOT NULL AUTO_INCREMENT, `account` varchar(50) NOT NULL DEFAULT '', `pwd` varchar(200) NOT NULL DEFAULT '', `name` varchar(50) NOT NULL DEFAULT '', `age` tinyint(4) DEFAULT NULL, PRIMARY KEY (`uid`), UNIQUE KEY `account` (`account`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4; ")
}

func TestDeleteUser(t *testing.T) {
	ResetDB()
	ulist := []User{
		{
			Name:    "mike",
			Account: "mike",
			Pwd:     "123",
			Age:     12,
		},
	}
	// create one
	CreateUser(&ulist[0])

	// check created
	u2, _ := GetUser(1)
	assert.Equal(t, &ulist[0], u2)

	DeleteUser(1)

	// check delete
	u2, _ = GetUser(1)
	assert.Nil(t, u2)
}

func TestCreateUser(t *testing.T) {
	ResetDB()
	ulist := []User{
		{
			Name:    "mike",
			Account: "mike",
			Pwd:     "123",
			Age:     12,
		},
		{
			Name:    "joe",
			Account: "joe",
			Pwd:     "321",
			Age:     24,
		},
	}

	idlist := make(map[int]User, len(ulist))
	for i := range ulist {
		x, _ := CreateUser(&ulist[i])
		idlist[x] = ulist[i]
	}

	for uid, u := range idlist {
		u2, err := GetUser(uid)
		assert.Nil(t, err)
		if uid != 0 {
			assert.NotNil(t, u2)
			assert.Equal(t, &u, u2)
		}
	}
}
