package ports

import (
	"context"
	"github.com/rodrigoferra/ports-blog/internal/core/domain"
)

type CommentService interface {
	Create(ctx context.Context, dto domain.Post) (*domain.Post, error)
	Delete(ctx context.Context, id string) error
}
