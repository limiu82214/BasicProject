package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPinPong(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/pin")
	if err != nil {
		t.Fatalf("TestPinPong error:[%v]", err)
		t.FailNow()
	}
	defer resp.Body.Close()
	assert.Equal(t, 200, resp.StatusCode, "http status code")

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	type respPong struct {
		Message string `json:"message"`
	}
	var pong respPong
	err = json.Unmarshal(body, &pong)
	assert.Nil(t, err)
	// fmt.Printf("message: %v\n", pong.Message)
	assert.Equal(t, "pong", pong.Message)
}
