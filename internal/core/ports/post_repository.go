package ports

import (
	"context"
	"myapi-modules/internal/core/domain"
)

type PostRepository interface {
	Create(ctx context.Context, entity *domain.Post) error
	GetByID(ctx context.Context, id string) (*domain.Post, error)
	List(ctx context.Context) ([]domain.Post, error)
	Update(ctx context.Context, entity *domain.Post) error
	Delete(ctx context.Context, id string) error
}
