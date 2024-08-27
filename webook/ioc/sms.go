package ioc

import (
	"github.com/Teresajw/go_project/webook/internal/service/sms"
	"github.com/Teresajw/go_project/webook/internal/service/sms/memory"
)

func InitSMSService() sms.Service {
	// 这里可以换其他的，暂时用内存实现
	return memory.NewService()
}
