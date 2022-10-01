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
