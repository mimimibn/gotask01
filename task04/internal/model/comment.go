package model

type Comment struct {
	// id 、 content 、 user_id （关联 users 表的 id ）、 post_id （关联 posts 表的 id ）、 created_at 等字段。
	BaseEntity
	Content string `gorm:"column:content;type:varchar(20);default:NULL;" json:"content"`
	UserId  int32
	PostId  int32
}

type CommentAddDTO struct {
	PostId  int32  `json:"postId" binding:"required"`
	Content string `json:"content" binding:"required"`
}
