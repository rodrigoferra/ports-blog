package dtos

type PostCreateDTO struct {
	Subject string
	Content string
}

type PostUpdateDTO struct {
	ID      string
	Subject *string
	Content *string
}
