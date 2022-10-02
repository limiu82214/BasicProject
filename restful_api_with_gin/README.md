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
v0.0.4 => 使用POST新增使用者資料 `/user`，使用leveldb來存放user資料  

* leveldb使用db/member/來儲存會員資料
* leveldb 是key value的儲存形式，所以沒辦法直接放入struct，先只儲存姓名
* 新增 CreateUser 用於新增使用者，並讓接手POST呼叫
* *todo* 測試前清空db
* *todo* 測試不應該與正式db混在一起
* ~~*todo* leveldb應該要使用單例模式~~

v0.0.5 => 使用gob來儲存user的資料 (將struct放入db、將struct序列化)

* *todo* 改單例後close成了問題，要實作連接池

v0.0.6 => db改成單例模式，使用signal的方式讓伺服器關閉前可以斷開db連接
v0.0.7 => DELETE 刪除資料 `/user/*`


## RESTful API

* `POST` `user` // data 已json的形式帶入 post body
* `GET` `user/*`
* `PUT`
* `PATH`
* `DELETE` `user/*`

### return

* 200 ok `v0.0.1`
* 201 created `v0.0.4`
* 202 accepted
* 204 no content
* 400 bad request `v0.0.4`
* 401 unauthorized
* 403 forbidden
* 404 not found
* 410 gone
* 500 internal server error `v0.0.4`
* 502 bad gateway
* 503 service unavailable
* 504 gateway timeout

## Point

* http request 的Test寫法 v0.0.1
* leveldb的使用 v0.0.4
* gob將struct序列化 v0.0.5
* leveldb單例封裝 v0.0.6
* 監聽signal，db連接失敗或重大錯誤時利用signal關閉server v0.0.6
