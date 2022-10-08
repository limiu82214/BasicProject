package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/limiu82214/GoBasicProject/restful_api_with_gin/db/user"
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

// TestPinPong 測試伺服器有無回應
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

// TestPostUser 測試新增user
func TestPostUser(t *testing.T) {
	ResetDB()
	uri := "http://localhost:8080/user"
	ulist := []user.User{
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

	idlist := make(map[int]user.User, len(ulist))
	for i, u := range ulist {
		b, _ := json.Marshal(u)
		resp, err := http.PostForm(uri, url.Values{"user": []string{string(b)}, "uid": []string{strconv.Itoa(i + 1)}})
		assert.Nil(t, err)
		defer resp.Body.Close()
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		uid, _ := ioutil.ReadAll(resp.Body)
		t_i, _ := strconv.Atoi(string(uid))
		ulist[i].Uid = t_i
		idlist[t_i] = ulist[i]
	}

	{ // read
		for uid, u := range idlist {
			resp, err := http.Get(uri + "/" + strconv.Itoa(uid))
			assert.Nil(t, err)
			defer resp.Body.Close()
			assert.Equal(t, 200, resp.StatusCode)

			body, err := ioutil.ReadAll(resp.Body)
			assert.Nil(t, err)
			j, _ := json.Marshal(u)
			assert.Equal(t, string(j), string(body))
		}
	}
}

// TestPostUser 測試刪除user
func TestDeleteUser(t *testing.T) {
	ResetDB()
	uri := "http://localhost:8080/user"
	u := &user.User{
		Name:    "mike",
		Account: "mike",
		Pwd:     "123",
		Age:     12,
	}
	b, _ := json.Marshal(u)

	// create one
	resp, err := http.PostForm(uri, url.Values{"user": []string{string(b)}, "uid": []string{`1`}})
	assert.Nil(t, err)
	defer resp.Body.Close()
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	uid, _ := ioutil.ReadAll(resp.Body)
	u.Uid, _ = strconv.Atoi(string(uid))

	// check created
	resp, err = http.Get(uri + "/1")
	assert.Nil(t, err)
	defer resp.Body.Close()
	assert.Equal(t, 200, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	j, _ := json.Marshal(u)
	assert.Equal(t, string(j), string(body))

	// delete one
	req, err := http.NewRequest(http.MethodDelete, uri+"/1", nil)
	assert.Nil(t, err)
	_, err = http.DefaultClient.Do(req)
	assert.Nil(t, err)

	time.Sleep(time.Second * 5) // for cache
	// check deleted
	resp, err = http.Get(uri + "/1")
	assert.Nil(t, err)
	defer resp.Body.Close()
	assert.Equal(t, 200, resp.StatusCode)

	body, err = ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, `null`, string(body))
}
