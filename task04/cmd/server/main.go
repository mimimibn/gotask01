package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"task04/internal/config"
	"task04/internal/controller"
	"task04/internal/middleware"
	"task04/internal/repository"
)

func main() {
	//用来练习读取配置文件，只加载了mysql.dns
	config, configErr := config.LoadConfig("./internal/config/dev.yaml")
	if configErr != nil {
		log.Fatalf("加载配置失败: %v", configErr)
	}
	r := gin.Default()
	if err := repository.InitDB(config.Mysql.Dns); err != nil {
		log.Fatal("数据库初始化失败:", err)
	}
	api := r.Group("/api")
	{
		authGroup := api.Group("/auth")
		{
			authGroup.POST("/register", controller.Register)
			authGroup.POST("/login", controller.Login)
		}
		postGroup := api.Group("/post")
		{
			//这里加中间件
			postGroup.Use(middleware.CheckJWT())
			postGroup.POST("/createPost", controller.CreatePost)
			postGroup.GET("/findList", controller.FindListPost)
			postGroup.GET("/findById", controller.FindByIdPost)
			postGroup.POST("/updateById", controller.UpdateByIdPost)
			postGroup.DELETE("/deleteById/:id", controller.DeleteByIdPost)
		}
		commentGroup := api.Group("/comment")
		{
			//这里加中间件
			commentGroup.Use(middleware.CheckJWT())
			commentGroup.POST("/createComment", controller.CreateComment)
			commentGroup.GET("/findCommentListByPostId", controller.FindCommentListByPostId)
		}
	}
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}
