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
	r.GET("/user/1", func(ctx *gin.Context) {
		u := db.NewUser()
		u.Name = "mike"
		ctx.JSON(http.StatusOK, u)
	})
	r.Run(":8080")
}
