package service

import (
	"context"
	"errors"
	"github.com/Teresajw/go_project/webook/internal/domain"
	"github.com/Teresajw/go_project/webook/internal/repository"
	repomocks "github.com/Teresajw/go_project/webook/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_userService_Login(t *testing.T) {
	testCases := []struct {
		name string
		mock func(ctrl *gomock.Controller) repository.UserRepository
		//输入
		ctx      context.Context
		email    string
		password string

		//期望
		wantUser domain.User
		wantErr  error
	}{
		{
			name: "正常登录",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				userRepo := repomocks.NewMockUserRepository(ctrl)
				// 模拟查询结果
				userRepo.EXPECT().FindByEmail(gomock.Any(), "2534447245@qq.com").Return(domain.User{
					Email:    "2534447245@qq.com",
					Password: "$2a$10$FFXYQdP2tzyWZldUlRRQo.bL724KfICxCjsmUIIIGSkGtzo5ZrmJG",
				}, nil)
				return userRepo
			},
			ctx:      context.Background(),
			email:    "2534447245@qq.com",
			password: "Abc@1234!",
			wantUser: domain.User{
				Email:    "2534447245@qq.com",
				Password: "$2a$10$FFXYQdP2tzyWZldUlRRQo.bL724KfICxCjsmUIIIGSkGtzo5ZrmJG",
			},
			wantErr: nil,
		},
		{
			name: "用户不存在",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				userRepo := repomocks.NewMockUserRepository(ctrl)
				// 模拟查询结果
				userRepo.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).Return(domain.User{}, repository.ErrUserNotFound)
				return userRepo
			},
			ctx:      context.Background(),
			email:    "2534447245@qq.com",
			password: "Abc@1234!",
			wantUser: domain.User{},
			wantErr:  ErrInvalidUserOrPassword,
		},
		{
			name: "密码错误",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				userRepo := repomocks.NewMockUserRepository(ctrl)
				// 模拟查询结果
				userRepo.EXPECT().FindByEmail(gomock.Any(), "2534447245@qq.com").Return(domain.User{
					Email:    "2534447245@qq.com",
					Password: "$2a$10$FFXYQdP2tzyWZldUlRRQo.bL724KfICxCjsmUIIIGSkGtzo5ZrmJG",
				}, nil)
				return userRepo
			},
			ctx:      context.Background(),
			email:    "2534447245@qq.com",
			password: "Abc@1234!1",
			wantUser: domain.User{},
			wantErr:  ErrInvalidUserOrPassword,
		},
		{
			name: "系统错误",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				userRepo := repomocks.NewMockUserRepository(ctrl)
				// 模拟查询结果
				userRepo.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).Return(domain.User{}, errors.New("mock db error"))
				return userRepo
			},
			ctx:      context.Background(),
			email:    "2534447245@qq.com",
			password: "Abc@1234!",
			wantUser: domain.User{},
			wantErr:  errors.New("mock db error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 准备
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			// 具体的测试代码
			userSvc := NewUserService(tc.mock(ctrl))
			user, err := userSvc.Login(tc.ctx, tc.email, tc.password)
			assert.Equal(t, tc.wantUser, user)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
