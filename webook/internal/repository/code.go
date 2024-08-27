package repository

import (
	"context"
	"github.com/Teresajw/go_project/webook/internal/repository/cache"
)

var (
	ErrCodeSendTooMany   = cache.ErrCodeSendTooMany
	ErrCodeVerifyTooMany = cache.ErrCodeVerifyTooMany
)

type CodeRepository struct {
	cache *cache.CodeCache
}

func NewCodeRepository(cache *cache.CodeCache) *CodeRepository {
	return &CodeRepository{cache: cache}
}

func (r *CodeRepository) Store(ctx context.Context, biz, phone, code string) error {
	return r.cache.Set(ctx, biz, phone, code)
}

func (r *CodeRepository) Verify(ctx context.Context, biz, phone, inputCode string) (bool, error) {
	return r.cache.Verify(ctx, biz, phone, inputCode)
}
