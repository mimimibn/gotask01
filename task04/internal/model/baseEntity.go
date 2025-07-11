package model

import "time"

type BaseEntity struct {
	Id        int32     `gorm:"column:id;type:int(11);primaryKey;not null;" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:NULL;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;default:NULL;" json:"updated_at"`
}
