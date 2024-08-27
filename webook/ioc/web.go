package ioc

import (
	"github.com/Teresajw/go_project/webook/internal/web"
	"github.com/Teresajw/go_project/webook/internal/web/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func InitGin(mdls []gin.HandlerFunc, dhl *web.UserHandler) *gin.Engine {
	server := gin.Default()
	server.Use(mdls...)
	dhl.RegisterRouters(server)
	return server
}

func InitMiddlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		cors.New(cors.Config{
			AllowOrigins: []string{"http://localhost:3000"},
			AllowMethods: []string{"POST", "OPTIONS", "GET", "PUT", "DELETE"},
			AllowHeaders: []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
			//不加这个前端拿不到,允许取的字段
			ExposeHeaders:    []string{"x-jwt-token"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				if strings.HasPrefix(origin, "http://localhost") {
					return true
				}
				return strings.Contains(origin, "https://localhost")
			},
			MaxAge: 12 * time.Hour,
		}),
		middleware.NewLoginJWTMiddlewareBuilder().IgnorePaths(
			"/users/signup",
			"/users/login",
			"/users/login_sms",
			"/users/login_sms/code/send").Build(),
	}
}
