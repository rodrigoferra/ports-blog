package adapters

import (
	"context"
	"myapi-modules/internal/core/domain"
	"myapi-modules/internal/core/ports"
	"myapi-modules/internal/infrastructure/dtos"
	"time"
)

type PostServiceImpl struct {
	repo ports.PostRepository
}

func NewPostService(repo ports.PostRepository) ports.PostService {
	return &PostServiceImpl{repo: repo}
}

func (p *PostServiceImpl) Create(ctx context.Context, dto dtos.PostCreateDTO) (*domain.Post, error) {
	post := domain.CreatePostFromDto(dto)
	if err := p.repo.Create(ctx, post); err != nil {
		return nil, err
	}
	return post, nil
}

func (p *PostServiceImpl) GetByID(ctx context.Context, id string) (*domain.Post, error) {
	return p.repo.GetByID(ctx, id)
}

func (p *PostServiceImpl) List(ctx context.Context) ([]domain.Post, error) {
	return p.repo.List(ctx)
}

func (p *PostServiceImpl) Update(ctx context.Context, id string, dto dtos.PostUpdateDTO) error {
	post, err := p.repo.GetByID(ctx, id)
	if err != nil || post == nil {
		return err
	}
	if dto.Content != nil {
		post.Content = *dto.Content
	}
	if dto.Subject != nil {
		post.Subject = *dto.Subject
	}
	post.UpdatedAt = time.Now()
	return p.repo.Update(ctx, post)
}

func (p *PostServiceImpl) Delete(ctx context.Context, id string) error {
	return p.repo.Delete(ctx, id)
}
