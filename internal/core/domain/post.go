package domain

import (
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

type PostCreateDTO struct {
	Subject string
	Content string
}

type PostUpdateDTO struct {
	ID      string
	Subject *string
	Content *string
}
