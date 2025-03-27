package entities

import "time"

type PostEntity struct {
	ID        string          `gorm:"primaryKey;type:varchar(255);not null" json:"id"`
	Subject   string          `gorm:"type:varchar(255);not null" json:"subject"`
	Content   string          `gorm:"type:text;not null" json:"content"`
	CreatedAt time.Time       `json:"created_at;not null"`
	UpdatedAt time.Time       `json:"updated_at"`
	Comments  []CommentEntity `gorm:"foreignKey:PostID" json:"comments"`
}

func (PostEntity) TableName() string {
	return "posts"
}
