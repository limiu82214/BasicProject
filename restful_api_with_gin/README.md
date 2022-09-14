# 使用gin實現的RESTful API

github.com/limiu82214/GoBasicProject/restful_api_with_gin

使用gin完成restful_api
[github gin](https://github.com/gin-gonic/gin)

## 規劃

* 接收與返回資料格式皆為JSON
* `使用者`新增查詢等功能

## Version

v0.0.1 => 實現了pin/pong呼叫  
v0.0.2 => 使用GET取回使用者資料 `/user/1`  
v0.0.3 => 使用GET取回使用者資料，可以應對不存在的ID `/user/*`  
v0.0.4 => 使用POST新增使用者資料 `/user`

## RESTful API

* `POST`
* `GET` `user/*`
* `PUT`
* `PATH`
* `DELETE`

### return

* 200 ok `v0.0.1`
* 201 created `v0.0.4`
* 202 accepted
* 204 no content
* 400 bad request
* 401 unauthorized
* 403 forbidden
* 404 not found
* 410 gone
* 500 internal server error
* 502 bad gateway
* 503 service unavailable
* 504 gateway timeout

# Point

* http request 的Test寫法 v0.0.1
* leveldb的使用 v0.0.4
