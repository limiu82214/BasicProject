package myutil

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/limiu82214/GoBasicProject/restful_api_with_gin/myutil/myredis"
)

func DefaultDecorator(h gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h(ctx)
	}
}

type HandlerFuncWithResult func(*gin.Context) (any, error)

func CacheDecorator(h HandlerFuncWithResult, param string, redisKeyPattern string,
	empty interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param(param)
		redisKey := fmt.Sprintf(redisKeyPattern, id)
		conn := myredis.RedisDefaultPool.Get()
		defer conn.Close()
		rst, err := redis.Bytes(conn.Do("get", redisKey))
		if err != nil {
			rst, err := h(ctx)
			if err != nil {
				ctx.JSON(rst.(int), err)
			}
			j, _ := json.Marshal(rst)
			conn.Do("setex", redisKey, 5, j)
			ctx.JSON(http.StatusOK, rst)
		} else {
			ctx.JSON(http.StatusOK, string(rst))
		}
	}
}
