1.将search.postman_collection.json导入postman
2.go环境 1.23.10
安装环境：
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
启动入口cmd/server/main.go
默认端口8080