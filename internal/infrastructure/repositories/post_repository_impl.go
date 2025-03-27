package repositories

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"myapi-modules/internal/core/domain"
	"myapi-modules/internal/core/ports"
)

type PostRepositoryImpl struct {
	db *gorm.DB
}

func NewGormPostRepository(db *gorm.DB) ports.PostRepository {
	return &PostRepositoryImpl{db: db}
}

func (p *PostRepositoryImpl) Create(ctx context.Context, post *domain.Post) error {
	return p.db.WithContext(ctx).Create(post).Error
}

func (p *PostRepositoryImpl) GetByID(ctx context.Context, id string) (*domain.Post, error) {
	var post domain.Post
	result := p.db.WithContext(ctx).First(&post, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &post, nil
}

func (p *PostRepositoryImpl) List(ctx context.Context) ([]domain.Post, error) {
	var posts []domain.Post
	result := p.db.WithContext(ctx).Find(&posts)
	return posts, result.Error
}

func (p *PostRepositoryImpl) Update(ctx context.Context, post *domain.Post) error {
	return p.db.WithContext(ctx).Save(post).Error
}

func (p *PostRepositoryImpl) Delete(ctx context.Context, id string) error {
	return p.db.WithContext(ctx).Delete(&domain.Post{}, id).Error
}
