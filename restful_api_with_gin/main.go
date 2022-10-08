package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/limiu82214/GoBasicProject/restful_api_with_gin/db/user"
	"github.com/limiu82214/GoBasicProject/restful_api_with_gin/myutil"
	"github.com/limiu82214/GoBasicProject/restful_api_with_gin/myutil/db"
	"github.com/limiu82214/GoBasicProject/restful_api_with_gin/myutil/myredis"
	"github.com/limiu82214/GoBasicProject/restful_api_with_gin/myutil/sig"
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

		v2.GET("/:uid", myutil.CacheDecorator(user.DaoGetUser, "uid", "uid_%s", user.User{}))
		v2.POST("", myutil.DefaultDecorator(user.DaoPostUser))
		v2.DELETE("/:uid", myutil.DefaultDecorator(user.DaoDeleteUser))
	}

	go (func() {
		debugPtr := flag.Bool("production", false, "change path to production path")
		flag.Parse()
		if *debugPtr {
		} else {
			d := db.GetInst()
			d.LogMode(true)
		}
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

	sig.ServerNotify()
	log.Println("伺服器開始關閉...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	log.Println("DB正在斷開連接...")
	db.DisconnectDB()
	log.Println("Redis正在斷開連接...")
	myredis.RedisDefaultPool.Close()
	log.Println("伺服器正在關閉...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalln("伺服器錯誤退出:", err)
	}
	log.Println("伺服器正常運行結束")
}
