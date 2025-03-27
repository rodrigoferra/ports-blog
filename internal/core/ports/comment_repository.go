package ports

import (
	"context"
	"myapi-modules/internal/core/domain"
)

type CommentRepository interface {
	Create(ctx context.Context, entity *domain.Comment) error
	Delete(ctx context.Context, id string) error
}
