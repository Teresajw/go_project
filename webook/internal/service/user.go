package service

import (
	"context"
	"errors"
	"github.com/Teresajw/go_project/webook/internal/domain"
	"github.com/Teresajw/go_project/webook/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

// var ErrDuplicateEmail = fmt.Errorf("%w", repository.ErrDuplicateEmail)
var (
	ErrDuplicateEmail        = repository.ErrDuplicateEmail
	ErrUserNotFound          = repository.ErrUserNotFound
	ErrInvalidUserOrPassword = errors.New("邮箱或者密码错误")
)

type UserService struct {
	// TODO: 依赖注入
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

//func (svc *UserService) Edit(ctx context.Context, user domain.User) error {
//	s, err := svc.repo.FindById(ctx, user.Id)
//}

func (svc *UserService) Profile(ctx context.Context, id int64) (domain.User, error) {
	user, err := svc.repo.FindById(ctx, id)
	if errors.Is(err, ErrUserNotFound) {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (svc *UserService) Login(ctx context.Context, email, password string) (domain.User, error) {
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

func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {
	// TODO: 加密
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	// TODO: 保存到数据库
	return svc.repo.Create(ctx, u)
}
