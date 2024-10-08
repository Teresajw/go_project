package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Teresajw/go_project/webook/internal/domain"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	ErrorUserNotFound = redis.Nil
)

var _ UserCache = (*RedisUserCache)(nil)

type UserCache interface {
	Get(ctx context.Context, id int64) (domain.User, error)
	Set(ctx context.Context, user domain.User) error
}

// 面相接口
type RedisUserCache struct {
	//传单机redis可以
	// 传集群redis
	cmd        redis.Cmdable
	expiration time.Duration
}

// A用到了B,B一定是接口===>面相接口编程
// A用到了B,B一定是A的字段===>规避包变量，包方法
// A用到了B, A一定不初始化B,而是外面注入，依赖注入(Dependecy Injection) 依赖反转(Inversion of Control)
func NewUserCache(cmd redis.Cmdable, expiration time.Duration) UserCache {
	return &RedisUserCache{
		cmd:        cmd,
		expiration: expiration,
	}
}

func (uc *RedisUserCache) Get(ctx context.Context, id int64) (domain.User, error) {
	key := uc.key(id)
	var user domain.User
	// 获取缓存
	bytes, err := uc.cmd.Get(ctx, key).Bytes()
	if err != nil {
		return user, err
	}
	err = json.Unmarshal(bytes, &user)
	return user, err
}

func (uc *RedisUserCache) Set(ctx context.Context, user domain.User) error {
	val, err := json.Marshal(user)
	if err != nil {
		return err
	}
	key := uc.key(user.Id)
	return uc.cmd.Set(ctx, key, val, uc.expiration).Err()
}

func (uc *RedisUserCache) key(id int64) string {
	return fmt.Sprintf("user:info:%d", id)
}

type UnifyCache interface {
	Get(ctx context.Context, id int64) (domain.User, error)
	Set(ctx context.Context, id int64, val any, expiration time.Duration) error
}
