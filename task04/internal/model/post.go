package model

type Post struct {
	BaseEntity
	Title    string `gorm:"column:title;type:varchar(20);default:NULL;" json:"title"`
	Content  string `gorm:"column:content;type:varchar(20);default:NULL;" json:"content"`
	UserId   int32
	Comments []Comment `gorm:"foreignKey:PostId"`
}

type PostAddDTO struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}
type PostUpdDTO struct {
	Id      int32  `json:"id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}
