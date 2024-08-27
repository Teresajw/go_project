//go:build wireinject

package main

import (
	"github.com/Teresajw/go_project/webook/internal/repository"
	"github.com/Teresajw/go_project/webook/internal/repository/cache"
	"github.com/Teresajw/go_project/webook/internal/repository/dao"
	"github.com/Teresajw/go_project/webook/internal/service"
	"github.com/Teresajw/go_project/webook/internal/web"
	"github.com/Teresajw/go_project/webook/ioc"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"time"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		ProvideExpiration,
		//基础依赖
		ioc.InitDB, ioc.InitRedis,
		//dao层
		dao.NewUserDao,
		cache.NewUserCache,
		cache.NewCodeCache,

		//repo层
		repository.NewUserRepository,
		repository.NewCodeRepository,
		//service层
		service.NewUserService,
		service.NewCodeService,

		ioc.InitSMSService,

		web.NewUserHandler,
		//middleware
		ioc.InitMiddlewares,
		//gin
		ioc.InitGin,
	)
	return new(gin.Engine)
}

func ProvideExpiration() time.Duration {
	return 15 * time.Minute // 或者你想要的任何时长
}
