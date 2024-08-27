//go:build wireinject

// 让wire注入代码
package wire

import (
	"github.com/Teresajw/go_project/wire/repository"
	"github.com/Teresajw/go_project/wire/repository/dao"
	"github.com/google/wire"
)

func InitRepository() *repository.UserRepository {
	// 这个方法里面传入各个组件的初始化方法
	wire.Build(repository.NewUserRepository, dao.NewUserDao, InitDB)
	return new(repository.UserRepository)
}
