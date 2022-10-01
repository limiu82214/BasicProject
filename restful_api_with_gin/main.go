package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/limiu82214/GoBasicProject/restful_api_with_gin/db"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	r := gin.Default()
	r.GET("/pin", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/user/:uid", func(ctx *gin.Context) {
		suid := ctx.Param("uid")
		uid, err := strconv.Atoi(suid)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		}
		u, err := db.GetUser(uid)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}
		ctx.JSON(http.StatusOK, u)
	})
	r.POST("/user", func(ctx *gin.Context) {
		u := db.NewUser()
		uid, _ := strconv.Atoi(ctx.DefaultPostForm("uid", "0"))
		err := json.Unmarshal([]byte(ctx.PostForm("user")), u)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}
		db.CreateUser(uid, u)
		ctx.JSON(http.StatusCreated, uid)
	})

	r.Run(":8080")
}
