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
		assert.Equal(t, name, user.Name)
	}

}
