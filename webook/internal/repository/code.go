package repository

import (
	"context"
	"github.com/Teresajw/go_project/webook/internal/repository/cache"
)

var (
	ErrCodeSendTooMany   = cache.ErrCodeSendTooMany
	ErrCodeVerifyTooMany = cache.ErrCodeVerifyTooMany
)

var _ CodeRepository = (*codeRepository)(nil)

type CodeRepository interface {
	Store(ctx context.Context, biz, phone, code string) error
	Verify(ctx context.Context, biz, phone, inputCode string) (bool, error)
}

type codeRepository struct {
	cache cache.CodeCache
}

func NewCodeRepository(cache cache.CodeCache) CodeRepository {
	return &codeRepository{cache: cache}
}

func (r *codeRepository) Store(ctx context.Context, biz, phone, code string) error {
	return r.cache.Set(ctx, biz, phone, code)
}

func (r *codeRepository) Verify(ctx context.Context, biz, phone, inputCode string) (bool, error) {
	return r.cache.Verify(ctx, biz, phone, inputCode)
}
