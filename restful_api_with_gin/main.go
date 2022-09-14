package main

import (
	"net/http"

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
		u := db.NewUser()
		isUIDExist := true
		switch ctx.Param("uid") {
		case "1":
			u.Name = "mike" // 應在資料庫存取 *todo*
		case "2":
			u.Name = "joe"
		default:
			isUIDExist = false
		}
		if !isUIDExist {
			ctx.JSON(http.StatusOK, struct{}{})
		} else {
			ctx.JSON(http.StatusOK, u)

		}
	})
	r.POST("/user", func(ctx *gin.Context) {
		ctx.JSON(http.StatusCreated, struct{}{}) // 應返回新增user的sn *todo*
	})

	r.Run(":8080")
}
