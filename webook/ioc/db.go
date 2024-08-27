package ioc

import (
	"github.com/Teresajw/go_project/webook/config"
	"github.com/Teresajw/go_project/webook/internal/repository/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.Config.DB.Dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	err = dao.InitTables(db)
	if err != nil {
		panic("failed to init tables")
	}
	return db
}
