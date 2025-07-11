package model

type User struct {
	BaseEntity
	UserName string `gorm:"column:user_name;type:varchar(20);default:NULL;" json:"user_name"`
	Password string `gorm:"column:password;type:varchar(255);default:NULL;" json:"password"`
	Posts    []Post `gorm:"foreignKey:UserId"`
}
