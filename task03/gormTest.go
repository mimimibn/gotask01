package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。
题目2：关联查询
基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。
题目3：钩子函数
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/

type User struct {
	Id        int32
	Name      string
	PostCount int32
	Post      []Post `gorm:"foreignKey:UserId"`
}

type Post struct {
	Id      int32
	Text    string
	UserId  int32
	Status  string
	Comment []Comment `gorm:"foreignKey:PostId"`
}

type Comment struct {
	Id     int32
	Com    string
	PostId int32
}

func (p *Post) AfterCreate(db *gorm.DB) (err error) {
	var count int64
	db.Model(&Post{}).Where("user_id = ?", p.UserId).Count(&count)
	db.Model(&User{}).Where("id = ?", p.UserId).Update("post_count", count)
	return
}
func (c Comment) AfterDelete(db *gorm.DB) (err error) {
	//为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
	var count int64
	db.Model(&Comment{}).Where("post_id = ?", c.PostId).Count(&count)
	if count == 0 {
		db.Debug().Model(&Post{}).Where("id = ?", c.PostId).Update("status", "无评论")
	}
	return
}

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		fmt.Println("链接数据库失败", err)
		return
	}
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		fmt.Println("创建数据库表失败")
		return
	}
	users := []User{
		{
			Name: "张三",
			Post: []Post{
				{Text: "文章111",
					Comment: []Comment{
						{Com: "我是文章1的评论1"},
						{Com: "我是文章1的评论2"},
					},
				},
				{Text: "文章2222",
					Comment: []Comment{
						{Com: "我是文章2的评论1"},
					},
				},
			},
		}, {
			Name: "李四",
			Post: []Post{
				{
					Text: "李四的文章",
					Comment: []Comment{
						{Com: "李四文章评论1"},
						{Com: "李四文章评论2"},
						{Com: "李四文章评论3"},
					}},
			},
		}}
	err = db.CreateInBatches(users, len(users)).Error
	if err != nil {
		fmt.Println("创建用户数据失败", err)
		return
	}
	//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	user := users[0]
	var u User //重新定义一个变量，接收查询出来的值
	err = db.Model(&User{}).Preload("Post.Comment").Where("id = ?", user.Id).First(&u).Error
	if err != nil {
		fmt.Println("数据不存在", err)
		return
	}
	fmt.Println(u)
	var postaa Post
	////编写Go代码，使用Gorm查询评论数量最多的文章信息。
	db.Model(&Post{}).InnerJoins("JOIN comments ON posts.id = comments.post_id").Group("Text").Order("count(comments.id) DESC").
		Limit(1).First(&postaa)
	fmt.Println(postaa)
	//测试新增文章钩子，设置用户的文章数量
	p := Post{Text: "我是第三个文章", UserId: 1}
	db.Save(&p)
	//测试删除评论时，为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
	c := Comment{PostId: 2}
	db.Where("post_id = ?", c.PostId).Delete(&c)
	//想要触发AfterDelete，则Delete必须传值，不能传空引用
	//在gorm中where的方法必须要在update前面，否则后面的where不会执行
}
