package service

import (
	"context"
	"errors"
	"github.com/Teresajw/go_project/webook/internal/domain"
	"github.com/Teresajw/go_project/webook/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

const DefaultPassword = "123456"

// var ErrDuplicateEmail = fmt.Errorf("%w", repository.ErrDuplicateEmail)
var (
	ErrDuplicateEmail        = repository.ErrDuplicateEmail
	ErrUserNotFound          = repository.ErrUserNotFound
	ErrInvalidUserOrPassword = errors.New("邮箱或者密码错误")
)

var _ UserService = (*userService)(nil)

type UserService interface {
	Profile(ctx context.Context, id int64) (domain.User, error)
	Login(ctx context.Context, email, password string) (domain.User, error)
	SignUp(ctx context.Context, u domain.User) error
	FindOrCreate(ctx context.Context, phone string) (domain.User, error)
	UpdateNoSensitiveInfo(ctx context.Context, user domain.User) error
}

type userService struct {
	// TODO: 依赖注入
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (svc *userService) UpdateNoSensitiveInfo(ctx context.Context, user domain.User) error {
	return svc.repo.UpdateNoSensitiveInfo(ctx, user)
}

func (svc *userService) Profile(ctx context.Context, id int64) (domain.User, error) {
	user, err := svc.repo.FindById(ctx, id)
	if errors.Is(err, ErrUserNotFound) {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (svc *userService) Login(ctx context.Context, email, password string) (domain.User, error) {
	u, err := svc.repo.FindByEmail(ctx, email)
	if errors.Is(err, ErrUserNotFound) {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	if err != nil {
		return domain.User{}, err
	}
	// TODO: 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		// 打印日志
		return domain.User{}, ErrInvalidUserOrPassword
	}
	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
	}, nil
}

func (svc *userService) SignUp(ctx context.Context, u domain.User) error {
	// TODO: 加密
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	// TODO: 保存到数据库
	return svc.repo.Create(ctx, u)
}

func (svc *userService) FindOrCreate(ctx context.Context, phone string) (domain.User, error) {
	u, err := svc.repo.FindByPhone(ctx, phone)
	if err != nil && !errors.Is(err, ErrUserNotFound) {
		return u, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(DefaultPassword), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, err
	}
	u = domain.User{
		Phone:    phone,
		Password: string(hash),
	}
	err = svc.repo.Create(ctx, u)
	if err != nil {
		return u, err
	}
	// 这里会遇到主从延迟的问题
	return svc.repo.FindByPhone(ctx, phone)
}
