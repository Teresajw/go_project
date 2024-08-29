package repository

import (
	"context"
	"database/sql"
	"github.com/Teresajw/go_project/webook/internal/domain"
	"github.com/Teresajw/go_project/webook/internal/repository/cache"
	cachemocks "github.com/Teresajw/go_project/webook/internal/repository/cache/mocks"
	"github.com/Teresajw/go_project/webook/internal/repository/dao"
	daomocks "github.com/Teresajw/go_project/webook/internal/repository/dao/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func Test_userRepository_FindById(t *testing.T) {
	now := time.Now()
	// 去掉毫秒以外的
	now = time.UnixMilli(now.UnixMilli())
	testCases := []struct {
		name string
		mock func(ctrl *gomock.Controller) (cache.UserCache, dao.UserDao)
		ctx  context.Context
		id   int64
		want domain.User
		err  error
	}{
		{
			name: "缓存未命中，查询db成功",
			mock: func(ctrl *gomock.Controller) (cache.UserCache, dao.UserDao) {
				userCache := cachemocks.NewMockUserCache(ctrl)
				userCache.EXPECT().Get(gomock.Any(), int64(123)).Return(domain.User{}, cache.ErrorUserNotFound)
				userDao := daomocks.NewMockUserDao(ctrl)
				userDao.EXPECT().FindById(gomock.Any(), int64(123)).Return(dao.User{
					Id: 123,
					Email: sql.NullString{
						String: "2534447245@qq.com",
						Valid:  true,
					},
					Password: "string",
					Nickname: "string",
					Phone: sql.NullString{
						String: "15056622919",
						Valid:  true,
					},
					Birthday: "string",
					Profile:  "string",
					Ctime:    now.UnixMilli(),
					Utime:    now.UnixMilli(),
				}, nil)
				userCache.EXPECT().Set(gomock.Any(), domain.User{
					Id:       123,
					Email:    "2534447245@qq.com",
					Password: "string",
					Nickname: "string",
					Phone:    "15056622919",
					Birthday: "string",
					Profile:  "string",
					Ctime:    now,
					Utime:    now,
				}).Return(nil)
				return userCache, userDao
			},
			ctx: context.Background(),
			id:  123,
			want: domain.User{
				Id:       123,
				Email:    "2534447245@qq.com",
				Password: "string",
				Nickname: "string",
				Phone:    "15056622919",
				Birthday: "string",
				Profile:  "string",
				Ctime:    now,
				Utime:    now,
			},
			err: nil,
		},
		{
			name: "缓存命中",
			mock: func(ctrl *gomock.Controller) (cache.UserCache, dao.UserDao) {
				userCache := cachemocks.NewMockUserCache(ctrl)
				userCache.EXPECT().Get(gomock.Any(), int64(123)).Return(domain.User{
					Id:       123,
					Email:    "2534447245@qq.com",
					Password: "string",
					Nickname: "string",
					Phone:    "15056622919",
					Birthday: "string",
					Profile:  "string",
					Ctime:    now,
					Utime:    now,
				}, nil)
				userDao := daomocks.NewMockUserDao(ctrl)
				return userCache, userDao
			},
			ctx: context.Background(),
			id:  123,
			want: domain.User{
				Id:       123,
				Email:    "2534447245@qq.com",
				Password: "string",
				Nickname: "string",
				Phone:    "15056622919",
				Birthday: "string",
				Profile:  "string",
				Ctime:    now,
				Utime:    now,
			},
			err: nil,
		},
		{
			name: "缓存未命中，查询db失败",
			mock: func(ctrl *gomock.Controller) (cache.UserCache, dao.UserDao) {
				userCache := cachemocks.NewMockUserCache(ctrl)
				userCache.EXPECT().Get(gomock.Any(), int64(123)).Return(domain.User{}, cache.ErrorUserNotFound)
				userDao := daomocks.NewMockUserDao(ctrl)
				userDao.EXPECT().FindById(gomock.Any(), int64(123)).Return(dao.User{}, dao.ErrUserNotFound)
				return userCache, userDao
			},
			ctx:  context.Background(),
			id:   123,
			want: domain.User{},
			err:  ErrUserNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			uc, ud := tc.mock(ctrl)
			userRepo := NewUserRepository(ud, uc)
			user, err := userRepo.FindById(tc.ctx, tc.id)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.want, user)
			time.Sleep(time.Second)
		})
	}
}
