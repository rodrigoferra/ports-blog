package ports

import (
	"context"
	"myapi-modules/internal/core/domain"
)

type CommentService interface {
	Create(ctx context.Context, dto domain.Post) (*domain.Post, error)
	Delete(ctx context.Context, id string) error
}
