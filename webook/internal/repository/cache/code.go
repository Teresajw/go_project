package cache

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
)

/*
编译器编译的时候会把set_code.lua文件嵌入到二进制文件中,也就是给变量luaSetCode赋值
*/
//go:embed lua/set_code.lua
var luaSetCode string

//go:embed lua/verify_code.lua
var luaVerifyCode string

var (
	// 发送频率限制
	ErrCodeSendTooMany   = errors.New("send code too frequent")
	ErrKeyNotTtl         = errors.New("key not ttl")
	ErrCodeVerifyTooMany = errors.New("verify code too frequent")
	ErrUnknownError      = errors.New("unknown error")
)

var _ CodeCache = (*RedisCodeCache)(nil)

type CodeCache interface {
	Set(ctx context.Context, biz, phone, code string) error
	Verify(ctx context.Context, biz, phone, inputCode string) (bool, error)
}

type RedisCodeCache struct {
	//传单机redis可以
	// 传集群redis
	cmd redis.Cmdable
}

func NewCodeCache(cmd redis.Cmdable) CodeCache {
	return &RedisCodeCache{cmd: cmd}
}

func (c *RedisCodeCache) Set(ctx context.Context, biz, phone, code string) error {
	val, err := c.cmd.Eval(ctx, luaSetCode, []string{c.key(biz, phone)}, code).Int64()
	if err != nil {
		return err
	}
	switch val {
	case 0:
		// 成功
		return nil
	case -1:
		// 发送太频繁
		return ErrCodeSendTooMany
	case -2:
		// 系统错误
		return ErrKeyNotTtl
	default:
		return ErrUnknownError
	}
}

func (c *RedisCodeCache) key(biz, phone string) string {
	return fmt.Sprintf("phone_code:%s:%s", biz, phone)
}

func (c *RedisCodeCache) Verify(ctx context.Context, biz, phone, inputCode string) (bool, error) {
	val, err := c.cmd.Eval(ctx, luaVerifyCode, []string{c.key(biz, phone)}, inputCode).Int()
	if err != nil {
		return false, err
	}
	switch val {
	case 0:
		return true, nil
	case -1:
		return false, ErrCodeVerifyTooMany
	case -2:
		return false, nil
	default:
		return false, ErrUnknownError
	}
}
