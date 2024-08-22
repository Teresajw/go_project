package repository

import (
	"context"
	"errors"
	"github.com/Teresajw/go_project/webook/internal/domain"
	"github.com/Teresajw/go_project/webook/internal/repository/cache"
	"github.com/Teresajw/go_project/webook/internal/repository/dao"
	"time"
)

// var ErrDuplicateEmail = fmt.Errorf("%w, 邮箱冲突", dao.ErrDuplicateEmail)
var (
	ErrDuplicateEmail = dao.ErrDuplicateEmail
	ErrUserNotFound   = dao.ErrUserNotFound
)

type UserRepository struct {
	dao   *dao.UserDao
	cache *cache.UserCache
}

func NewUserRepository(dao *dao.UserDao, cache *cache.UserCache) *UserRepository {
	return &UserRepository{
		dao:   dao,
		cache: cache,
	}
}

func (r *UserRepository) FindById(ctx context.Context, id int64) (domain.User, error) {
	// 优先从cache里面找
	u, err := r.cache.Get(ctx, id)
	switch {
	case err == nil:
		return u, nil
	case errors.Is(err, cache.ErrorUserNotFound):
		// 去db里面找
		user, err := r.dao.FindById(ctx, id)
		if err != nil {
			return domain.User{}, err
		}
		// 使用time.Unix和纳秒偏移量来创建time.Time对象
		ct := time.Unix(user.Ctime/1000, (user.Ctime%1000)*1000000)
		ut := time.Unix(user.Utime/1000, (user.Ctime%1000)*1000000)

		u = domain.User{
			Id:       user.Id,
			Email:    user.Email,
			Nickname: user.Nickname,
			Phone:    user.Phone,
			Birthday: user.Birthday,
			Profile:  user.Profile,
			Ctime:    ct,
			Utime:    ut,
		}
		// 回写cache
		go func() {
			// 回写cache
			err = r.cache.Set(ctx, u)
			if err != nil {
				//return domain.User{}, err
				//写日志就行
			}
		}()
		return u, err
	default:
		// redis 异常，不去mysql 查询，保护DB
		return domain.User{}, err
	}
}

//这里怎么办?err = i0.EOF
// 要不要去数据库加载?
//看起来我不应该加载?
// 看起来我好像也要加载?

//选加载- 做好兜底，万- Redis 真的崩了，你要保护住你的数据库
// 我数据库限流呀!

// 选不加载- 用户体验差一点
// 缓存里面有数据
// 缓存里面没有数据
// 缓存出错了，你也不知道有没有数据

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	// TODO: 先从cache里面找
	// TODO: 如果没有就从db里面找
	u, err := r.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
	}, nil
	// TODO: 找到了回写cache
}

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	return r.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
		Nickname: u.Nickname,
		Phone:    u.Phone,
		Birthday: u.Birthday,
		Profile:  u.Profile,
	})
	// TODO: 创建成功后，需要回写cache,操作缓存
}
