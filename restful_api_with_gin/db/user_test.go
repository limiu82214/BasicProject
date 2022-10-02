package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteUser(t *testing.T) {
	// create one
	u := NewUser()
	u.Name = "mike"
	CreateUser(1, u)

	// check created
	u2, _ := GetUser(1)
	assert.Equal(t, u, u2)

	DeleteUser(1)

	// check delete
	u2, _ = GetUser(1)
	assert.Nil(t, u2)
}

func TestCreateUser(t *testing.T) {
	u := NewUser()
	u.Name = "mike"
	CreateUser(1, u)
	u.Name = "joe"
	CreateUser(2, u)
	TestGetUser(t)
}

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
