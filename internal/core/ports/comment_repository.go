package ports

import (
	"context"
	"github.com/rodrigoferra/ports-blog/internal/core/domain"
)

type CommentRepository interface {
	Create(ctx context.Context, entity *domain.Comment) error
	Delete(ctx context.Context, id string) error
}
