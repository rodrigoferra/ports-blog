package ports

import (
	"context"
	"github.com/rodrigoferra/ports-blog/internal/core/domain"
)

type PostService interface {
	Create(ctx context.Context, dto domain.PostCreateDTO) (*domain.Post, error)
	GetByID(ctx context.Context, id string) (*domain.Post, error)
	List(ctx context.Context) ([]domain.Post, error)
	Update(ctx context.Context, id string, dto domain.PostUpdateDTO) error
	Delete(ctx context.Context, id string) error
}
