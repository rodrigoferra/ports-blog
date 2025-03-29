package ports

import (
	"context"
	"myapi-modules/internal/core/domain"
	"myapi-modules/internal/infrastructure/dtos"
)

type PostService interface {
	Create(ctx context.Context, dto dtos.PostCreateDTO) (*domain.Post, error)
	GetByID(ctx context.Context, id string) (*domain.Post, error)
	List(ctx context.Context) ([]domain.Post, error)
	Update(ctx context.Context, id string, dto dtos.PostUpdateDTO) error
	Delete(ctx context.Context, id string) error
}
