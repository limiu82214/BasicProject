package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DaoGetUser(ctx *gin.Context) (any, error) {
	suid := ctx.Param("uid")
	uid, err := strconv.Atoi(suid)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return GetUser(uid)
}

func DaoPostUser(ctx *gin.Context) {
	u := &User{}
	err := json.Unmarshal([]byte(ctx.PostForm("user")), u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	uid, _ := CreateUser(u)
	if uid != 0 {
		ctx.JSON(http.StatusCreated, u.Uid)
	} else {
		ctx.JSON(http.StatusConflict, u.Uid)
	}
}

func DaoDeleteUser(ctx *gin.Context) {
	suid := ctx.Param("uid")
	uid, err := strconv.Atoi(suid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	err = DeleteUser(uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, uid)
}
