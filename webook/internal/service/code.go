package service

import (
	"context"
	"fmt"
	"github.com/Teresajw/go_project/webook/internal/repository"
	"github.com/Teresajw/go_project/webook/internal/service/sms"
	"math/rand"
)

var (
	ErrCodeSendTooMany   = repository.ErrCodeSendTooMany
	ErrCodeVerifyTooMany = repository.ErrCodeVerifyTooMany
)

const (
	TplId = "1548751"
)

type CodeService struct {
	repo       *repository.CodeRepository
	smsService sms.Service //这是短信实现的接口
}

func NewCodeService(repo *repository.CodeRepository, smsService sms.Service) *CodeService {
	return &CodeService{repo: repo, smsService: smsService}
}

func (svc *CodeService) Send(ctx context.Context, biz, phone string) error {
	// 生成验证码
	code := svc.generateCode()
	//塞进去redis
	err := svc.repo.Store(ctx, biz, phone, code)
	if err != nil {
		// 有问题
		return err
	}
	return svc.smsService.Send(ctx, TplId, []string{code}, phone)

}

func (svc *CodeService) generateCode() string {
	num := rand.Intn(1000000)
	//不够6位补0
	return fmt.Sprintf("%06d", num)
}

func (svc *CodeService) Verify(ctx context.Context, biz, phone, inputCode string) (bool, error) {
	return svc.repo.Verify(ctx, biz, phone, inputCode)
}
