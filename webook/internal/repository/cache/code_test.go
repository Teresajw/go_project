package cache

import (
	"context"
	"errors"
	"github.com/Teresajw/go_project/webook/internal/repository/cache/redismocks"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestRedisCodeCache_Set(t *testing.T) {
	testCases := []struct {
		name  string
		mock  func(ctrl *gomock.Controller) redis.Cmdable
		ctx   context.Context
		biz   string
		phone string
		code  string
		want  error
	}{
		{
			name: "验证码设置成功",
			mock: func(ctrl *gomock.Controller) redis.Cmdable {
				cmd := redismocks.NewMockCmdable(ctrl)
				rd := redis.NewCmd(context.Background())
				rd.SetVal(int64(0))
				cmd.EXPECT().Eval(gomock.Any(), luaSetCode, []string{"phone_code:login:15056622919"}, []any{"8888"}).Return(rd)
				return cmd
			},
			ctx:   context.Background(),
			biz:   "login",
			phone: "15056622919",
			code:  "8888",
			want:  nil,
		},
		{
			name: "redis错误",
			mock: func(ctrl *gomock.Controller) redis.Cmdable {
				cmd := redismocks.NewMockCmdable(ctrl)
				rd := redis.NewCmd(context.Background())
				rd.SetErr(errors.New("redis error"))
				cmd.EXPECT().Eval(gomock.Any(), luaSetCode, []string{"phone_code:login:15056622919"}, []any{"8888"}).Return(rd)
				return cmd
			},
			ctx:   context.Background(),
			biz:   "login",
			phone: "15056622919",
			code:  "8888",
			want:  errors.New("redis error"),
		},
		{
			name: "发送太频繁",
			mock: func(ctrl *gomock.Controller) redis.Cmdable {
				cmd := redismocks.NewMockCmdable(ctrl)
				rd := redis.NewCmd(context.Background())
				rd.SetVal(int64(-1))
				cmd.EXPECT().Eval(gomock.Any(), luaSetCode, []string{"phone_code:login:15056622919"}, []any{"8888"}).Return(rd)
				return cmd
			},
			ctx:   context.Background(),
			biz:   "login",
			phone: "15056622919",
			code:  "8888",
			want:  ErrCodeSendTooMany,
		},
		{
			name: "系统错误",
			mock: func(ctrl *gomock.Controller) redis.Cmdable {
				cmd := redismocks.NewMockCmdable(ctrl)
				rd := redis.NewCmd(context.Background())
				rd.SetVal(int64(10))
				cmd.EXPECT().Eval(gomock.Any(), luaSetCode, []string{"phone_code:login:15056622919"}, []any{"8888"}).Return(rd)
				return cmd
			},
			ctx:   context.Background(),
			biz:   "login",
			phone: "15056622919",
			code:  "8888",
			want:  ErrUnknownError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			codeCache := NewCodeCache(tc.mock(ctrl))
			err := codeCache.Set(tc.ctx, tc.biz, tc.phone, tc.code)
			assert.Equal(t, tc.want, err)
		})
	}
}
