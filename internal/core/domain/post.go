package domain

import (
	"github.com/google/uuid"
	"myapi-modules/internal/infrastructure/dtos"
	"time"
)

type Post struct {
	ID        string
	Subject   string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Comments  []Comment
}

func CreatePostFromDto(p dtos.PostCreateDTO) *Post {
	return &Post{
		ID:        uuid.NewString(),
		Subject:   p.Subject,
		Content:   p.Content,
		CreatedAt: time.Now(),
	}
}
