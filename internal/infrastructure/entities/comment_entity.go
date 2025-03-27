package entities

import "time"

type CommentEntity struct {
	ID        string     `gorm:"primaryKey;type:varchar(255);not null" json:"id"`
	Content   string     `gorm:"type:text;not null" json:"content"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	PostID    string     `gorm:"index:idx_post;type:varchar(255);not null" json:"post_id"`
	Post      PostEntity `gorm:"foreignKey:PostID" json:"post"`
}

func (CommentEntity) TableName() string {
	return "comments"
}
