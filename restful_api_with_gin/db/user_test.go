package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	excepted := map[int]string{
		0: "",
		1: "mike",
		2: "joe",
	}

	for uid, name := range excepted {
		user, err := GetUser(uid)
		assert.Nil(t, err)
		if uid != 0 {
			assert.Equal(t, name, user.Name)
		}
	}

}

func TestCreateUser(t *testing.T) {
	u := NewUser()
	u.Name = "mike"
	CreateUser(1, u)
	u.Name = "joe"
	CreateUser(2, u)
	TestGetUser(t)
}
