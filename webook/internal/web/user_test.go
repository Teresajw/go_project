package web

import (
	"bytes"
	"errors"
	"github.com/Teresajw/go_project/webook/internal/service"
	svcmocks "github.com/Teresajw/go_project/webook/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandler_SignUp(t *testing.T) {
	testCases := []struct {
		name     string
		mock     func(ctrl *gomock.Controller) service.UserService
		reqBody  string
		wantCode int
		wantBody string
	}{
		{
			name: "【参数不对,bind绑定失败】",
			mock: func(ctrl *gomock.Controller) service.UserService {
				userSvc := svcmocks.NewMockUserService(ctrl)
				return userSvc
			},
			reqBody: `{
    "email":"2534447245@qq.com",
    "password":"Abc@1234!",
    "confirmPassword":"Abc@1234!"`,
			wantCode: http.StatusBadRequest,
		},
		{
			name: "【邮箱格式错误】",
			mock: func(ctrl *gomock.Controller) service.UserService {
				userSvc := svcmocks.NewMockUserService(ctrl)
				return userSvc
			},
			reqBody: `{
    "email":"2534447245@qq",
    "password":"Abc@1234!",
    "confirmPassword":"Abc@1234!",
    "nickname":"",
    "phone":"",
    "birthday":"",
    "profile":""
}`,
			wantCode: http.StatusOK,
			wantBody: "邮箱格式错误",
		},
		{
			name: "【两次密码不一致】",
			mock: func(ctrl *gomock.Controller) service.UserService {
				userSvc := svcmocks.NewMockUserService(ctrl)
				return userSvc
			},
			reqBody: `{
    "email":"2534447245@qq.com",
    "password":"Abc@1234!",
    "confirmPassword":"111Abc@1234!",
    "nickname":"",
    "phone":"",
    "birthday":"",
    "profile":""
}`,
			wantCode: http.StatusOK,
			wantBody: "两次密码不一致",
		},
		{
			name: "【密码格式不正确】",
			mock: func(ctrl *gomock.Controller) service.UserService {
				userSvc := svcmocks.NewMockUserService(ctrl)
				return userSvc
			},
			reqBody: `{
    "email":"2534447245@qq.com",
    "password":"1234",
    "confirmPassword":"1234",
    "nickname":"",
    "phone":"",
    "birthday":"",
    "profile":""
}`,
			wantCode: http.StatusOK,
			wantBody: "密码必须大于8位、包含数字、特殊字符",
		},
		{
			name: "【注册成功】",
			mock: func(ctrl *gomock.Controller) service.UserService {
				userSvc := svcmocks.NewMockUserService(ctrl)
				userSvc.EXPECT().SignUp(gomock.Any(), gomock.Any()).Return(nil)
				return userSvc
			},
			reqBody: `{
    "email":"2534447245@qq.com",
    "password":"Abc@1234!",
    "confirmPassword":"Abc@1234!",
    "nickname":"",
    "phone":"",
    "birthday":"",
    "profile":""
}`,
			wantCode: http.StatusOK,
			wantBody: `{"code":200,"msg":"注册成功！","data":{"email":"2534447245@qq.com"}}`,
		},
		{
			name: "【邮箱冲突】",
			mock: func(ctrl *gomock.Controller) service.UserService {
				userSvc := svcmocks.NewMockUserService(ctrl)
				userSvc.EXPECT().SignUp(gomock.Any(), gomock.Any()).Return(service.ErrDuplicateEmail)
				return userSvc
			},
			reqBody: `{
    "email":"2534447245@qq.com",
    "password":"Abc@1234!",
    "confirmPassword":"Abc@1234!",
    "nickname":"",
    "phone":"",
    "birthday":"",
    "profile":""
}`,
			wantCode: http.StatusOK,
			wantBody: "邮箱已经存在",
		},
		{
			name: "【系统异常】",
			mock: func(ctrl *gomock.Controller) service.UserService {
				userSvc := svcmocks.NewMockUserService(ctrl)
				userSvc.EXPECT().SignUp(gomock.Any(), gomock.Any()).Return(errors.New("system error"))
				return userSvc
			},
			reqBody: `{
    "email":"2534447245@qq.com",
    "password":"Abc@1234!",
    "confirmPassword":"Abc@1234!",
    "nickname":"",
    "phone":"",
    "birthday":"",
    "profile":""
}`,
			wantCode: http.StatusInternalServerError,
			wantBody: "系统异常",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 准备
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			// 注册路由
			server := gin.Default()
			// 利用mock 构造UserHandler
			h := NewUserHandler(tc.mock(ctrl), nil)
			h.RegisterRouters(server)
			// 构造请求
			req := httptest.NewRequest(http.MethodPost, "/users/signup", bytes.NewBufferString(tc.reqBody))
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()
			// 执行
			server.ServeHTTP(resp, req)
			// 断言
			assert.Equal(t, tc.wantCode, resp.Code)
			assert.Equal(t, tc.wantBody, resp.Body.String())
		})
	}
}

/*func TestMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userSvc := svcmocks.NewMockUserService(ctrl)
	userSvc.EXPECT().SignUp(gomock.Any(), gomock.Any()).Return(errors.New("mock error"))

	err := userSvc.SignUp(context.Background(), domain.User{
		Email: "2131312",
	})

	t.Log(err)

}*/
