package wire

import (
	"fmt"
	"github.com/Teresajw/go_project/wire/repository"
	"github.com/Teresajw/go_project/wire/repository/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/go_project?charset=utf8&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	userDao := dao.NewUserDao(db)
	repo := repository.NewUserRepository(userDao)
	fmt.Println(repo)
	InitRepository()
}
