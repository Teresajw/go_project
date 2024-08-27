package tencent

import (
	"context"
	"fmt"
	"github.com/ecodeclub/ekit"
	"github.com/ecodeclub/ekit/slice"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

type Service struct {
	appId  *string
	sign   *string
	client *sms.Client
}

func NewService(client *sms.Client, appId, sign string) *Service {
	return &Service{
		appId:  ekit.ToPtr[string](appId),
		sign:   ekit.ToPtr[string](sign),
		client: client,
	}
}

func (s *Service) Send(ctx context.Context, tplId string, args []string, numbers ...string) error {
	req := sms.NewSendSmsRequest()
	req.SmsSdkAppId = s.appId
	req.SignName = s.sign
	req.TemplateId = ekit.ToPtr[string](tplId)
	req.PhoneNumberSet = s.toStringPtrSlice(numbers)
	req.TemplateParamSet = s.toStringPtrSlice(args)
	resp, err := s.client.SendSms(req)
	if err != nil {
		return err
	}
	for _, status := range resp.Response.SendStatusSet {
		if status.Code == nil || *(status.Code) != "Ok" {
			return fmt.Errorf("发送短信失败: %s, 原因: %s", *status.Code, *status.Message)
		}
	}
	return nil
}

func (s *Service) toStringPtrSlice(src []string) []*string {
	/*dst := make([]*string, len(src))
	for i, v := range src {
		dst[i] = &v
	}
	return dst*/
	return slice.Map[string, *string](src, func(idx int, v string) *string {
		return &v
	})
}
