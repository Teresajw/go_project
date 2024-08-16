package repository

import (
	"context"
	"github.com/Teresajw/go_project/webook/internal/domain"
	"github.com/Teresajw/go_project/webook/internal/repository/dao"
	"time"
)

// var ErrDuplicateEmail = fmt.Errorf("%w, 邮箱冲突", dao.ErrDuplicateEmail)
var (
	ErrDuplicateEmail = dao.ErrDuplicateEmail
	ErrUserNotFound   = dao.ErrUserNotFound
)

type UserRepository struct {
	dao *dao.UserDao
}

func NewUserRepository(dao *dao.UserDao) *UserRepository {
	return &UserRepository{dao: dao}
}

func (r *UserRepository) FindById(ctx context.Context, id int64) (domain.User, error) {
	u, err := r.dao.FindById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	// 使用time.Unix和纳秒偏移量来创建time.Time对象
	ct := time.Unix(u.Ctime/1000, (u.Ctime%1000)*1000000)
	ut := time.Unix(u.Utime/1000, (u.Ctime%1000)*1000000)

	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Nickname: u.Nickname,
		Phone:    u.Phone,
		Birthday: u.Birthday,
		Profile:  u.Profile,
		Ctime:    ct,
		Utime:    ut,
	}, nil
	// TODO: 找到了回写cache
}

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
