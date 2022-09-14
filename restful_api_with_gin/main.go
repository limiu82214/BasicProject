package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/limiu82214/GoBasicProject/restful_api_with_gin/db"
)

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
		// 應去資料庫新增 user *todo*
		ctx.JSON(http.StatusCreated, struct{}{}) // 應返回新增user的sn *todo*
	})

	r.Run(":8080")
}
