package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"task04/internal/model"
)

var DB *gorm.DB

func InitDB(dns string) error {
	var err error
	//DB, err = gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"))
	DB, err = gorm.Open(mysql.Open(dns))
	if err != nil {
		return err
	}
	log.Println("链接数据库成功")
	//, &model.Post{}, &model.Comment{}
	err = DB.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
	log.Println("数据库表创建成功")
	return nil
}
