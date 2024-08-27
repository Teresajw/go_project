package tencent

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"os"
	"testing"
)

func TestService_Send(t *testing.T) {
	secretid, ok := os.LookupEnv("SecretId")
	if !ok {
		t.Fatal("SecretId not found")
	}
	secretkey, ok := os.LookupEnv("SecretKey")
	if !ok {
		t.Fatal("SecretKey not found")
	}
	// 实例化一个认证对象，入参需要传入腾讯云账户 SecretId，SecretKey
	credential := common.NewCredential(
		secretid,
		secretkey,
	)

	c, err := sms.NewClient(credential, "ap-guangzhou", profile.NewClientProfile())
	if err != nil {
		t.Fatal(err)
	}
	s := NewService(c, "", "")

	testCases := []struct {
		name    string
		tplId   string
		args    []string
		numbers []string
		wantErr error
	}{
		{
			name:  "发送验证码",
			tplId: "",
			args:  []string{"123456"},
			numbers: []string{
				"15056622919",
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			er := s.Send(context.Background(), tt.tplId, tt.args, tt.numbers...)
			assert.Equal(t, tt.wantErr, er)
		})
	}
}
