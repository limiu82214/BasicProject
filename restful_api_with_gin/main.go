package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/limiu82214/GoBasicProject/restful_api_with_gin/db"
	"github.com/limiu82214/GoBasicProject/restful_api_with_gin/myutil"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	r := gin.Default()
	v1 := r.Group("/")
	{
		v1.GET("pin", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	v2 := r.Group("/user")
	{
		v2.GET("/:uid", func(ctx *gin.Context) {
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

		v2.POST("", func(ctx *gin.Context) {
			u := db.NewUser()
			uid, _ := strconv.Atoi(ctx.DefaultPostForm("uid", "0"))
			err := json.Unmarshal([]byte(ctx.PostForm("user")), u)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
			}
			db.CreateUser(uid, u)
			ctx.JSON(http.StatusCreated, uid)
		})

		v2.DELETE("/:uid", func(ctx *gin.Context) {
			suid := ctx.Param("uid")
			uid, err := strconv.Atoi(suid)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, err)
			}
			err = db.DeleteUser(uid)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
			}
			ctx.JSON(http.StatusOK, uid)
		})
	}

	go (func() {
		myutil.GetInst()
	})()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	go (func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
		r.Run(":8080")
	})()

	myutil.ServerNotify()
	log.Println("伺服器開始關閉...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	log.Println("DB正在斷開連接...")
	myutil.DisconnectDB()
	log.Println("伺服器正在關閉...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalln("伺服器錯誤退出:", err)
	}
	log.Println("伺服器正常運行結束")
}
