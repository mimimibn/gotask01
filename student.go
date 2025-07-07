package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
*
题目1：基本CRUD操作
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

建表语句：
CREATE TABLE `students` (

	`id` int(11) NOT NULL AUTO_INCREMENT,
	`name` varchar(64) DEFAULT NULL,
	`age` int(11) DEFAULT NULL,
	`grade` varchar(64) DEFAULT NULL,
	PRIMARY KEY (`id`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

工具自动生成的gorm对应的结构体

	type Students struct {
		Id int32 `gorm:"column:id;type:int(11);primaryKey;not null;" json:"id"`
		Name string `gorm:"column:name;type:varchar(64);default:NULL;" json:"name"`
		Age int32 `gorm:"column:age;type:int(11);default:NULL;" json:"age"`
		Grade string `gorm:"column:grade;type:varchar(64);default:NULL;" json:"grade"`
	}
*/
type Students struct {
	Id    int32  `gorm:"column:id;type:int(11);primaryKey;not null;" json:"id"`
	Name  string `gorm:"column:name;type:varchar(64);default:NULL;" json:"name"`
	Age   int32  `gorm:"column:age;type:int(11);default:NULL;" json:"age"`
	Grade string `gorm:"column:grade;type:varchar(64);default:NULL;" json:"grade"`
}

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		fmt.Println("链接数据库失败", err)
		return
	}
	//编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	//student := Students{Name: "张三", Age: 20, Grade: "三年级"}
	//db.Save(&student)
	//编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	var student Students
	db.Model(&Students{}).Where(" age > ?", 18).Find(&student)
	fmt.Println(student)
	//编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	student.Grade = "四年级"
	db.Save(&student)
	//编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	db.Where("age < ?", 15).Delete(&Students{})
}
