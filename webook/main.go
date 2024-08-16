package main

import (
	"github.com/Teresajw/go_project/webook/internal/repository"
	"github.com/Teresajw/go_project/webook/internal/repository/dao"
	"github.com/Teresajw/go_project/webook/internal/service"
	"github.com/Teresajw/go_project/webook/internal/web"
	"github.com/Teresajw/go_project/webook/internal/web/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
	"time"
)

func main() {
	db := initDB()
	u := initUser(db)
	server := initWebServer()
	u.RegisterRouters(server)
	server.Run(":8080")
}

func initWebServer() *gin.Engine {
	server := gin.Default()
	/*server.Use(func(ctx *gin.Context) {
		print("这是第一个middleware")
	})
	server.Use(func(ctx *gin.Context) {
		print("这是第二个middleware")
	})*/
	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"POST", "OPTIONS", "GET", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		//ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "https://localhost")
		},
		MaxAge: 12 * time.Hour,
	}))
	//设置session
	//store := cookie.NewStore([]byte("secret"))
	//store := memstore.NewStore(
	//	[]byte("SUwcr3HfInY49a4XVQ03lV4u1AgcQkynTkf5dPbEAknqr8K7zh5WFNLLPgpUocHi"),
	//	[]byte("G4Pflse7D0753X6UK7jKAi4k9bGvHF1OkhtJzp9O5VRxgOoc58OBV4zbmFkyYTvL"),
	//)
	store, err := redis.NewStore(
		16,
		"tcp",
		"192.168.112.44:32738",
		"",
		[]byte("nXVOTQA7gXaQk1sroNkyQfDNYg6GagJF"),
		[]byte("vBg72COgIEfnTlvFvrXXGDJgcBVQ7kAe"),
	)
	if err != nil {
		panic("failed to connect redis")
	}

	server.Use(sessions.Sessions("mySession", store))
	server.Use(middleware.NewLoginMiddlewareBuilder().IgnorePaths("/users/signup", "/users/login").Build())
	return server
}

func initUser(db *gorm.DB) *web.UserHandler {
	ud := dao.NewUserDao(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	u := web.NewUserHandler(svc)
	return u
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:CQGWiRshWb@tcp(192.168.112.24:3306)/test?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{
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
