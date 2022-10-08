package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/limiu82214/GoBasicProject/restful_api_with_gin/myutil"
)

func DaoGetUser(ctx *gin.Context) (any, myutil.StatusErrorer) {
	suid := ctx.Param("uid")
	uid, err := strconv.Atoi(suid)
	if err != nil {
		return nil, myutil.NewStatusError(http.StatusBadRequest, err)
	}
	u, err := GetUser(uid)
	return u, myutil.NewStatusError(http.StatusOK, err)
}

func DaoPostUser(ctx *gin.Context) (any, myutil.StatusErrorer) {
	u := &User{}
	err := json.Unmarshal([]byte(ctx.PostForm("user")), u)
	if err != nil {
		return nil, myutil.NewStatusError(http.StatusInternalServerError, err)
	}
	uid, _ := CreateUser(u)
	if uid != 0 {
		return u.Uid, myutil.NewStatusErrorString(http.StatusCreated, "")
	} else {
		return u.Uid, myutil.NewStatusErrorString(http.StatusConflict, "existed")
	}
}

func DaoDeleteUser(ctx *gin.Context) (any, myutil.StatusErrorer) {
	suid := ctx.Param("uid")
	uid, err := strconv.Atoi(suid)
	if err != nil {
		return nil, myutil.NewStatusError(http.StatusBadRequest, err)
	}
	err = DeleteUser(uid)
	if err != nil {
		return nil, myutil.NewStatusErrorString(http.StatusInternalServerError, "StatusInternalServerError")
	}
	return uid, myutil.NewStatusErrorString(http.StatusOK, "")
}
