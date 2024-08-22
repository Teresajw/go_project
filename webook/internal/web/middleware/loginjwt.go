package middleware

import (
	"encoding/gob"
	"github.com/Teresajw/go_project/webook/internal/web"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

type LoginJWTMiddlewareBuilder struct {
	paths []string
}

func NewLoginJWTMiddlewareBuilder() *LoginJWTMiddlewareBuilder {
	return &LoginJWTMiddlewareBuilder{}
}

func (l *LoginJWTMiddlewareBuilder) IgnorePaths(paths ...string) *LoginJWTMiddlewareBuilder {
	for _, path := range paths {
		l.paths = append(l.paths, path)
	}
	return l
}

func (l *LoginJWTMiddlewareBuilder) Build() gin.HandlerFunc {
	gob.Register(time.Time{})
	return func(ctx *gin.Context) {
		for _, path := range l.paths {
			if ctx.Request.URL.Path == path {
				return
			}
		}
		// 使用JWT校验
		// 1. 获取token
		tokenHeader := ctx.GetHeader("Authorization")
		segs := strings.Split(tokenHeader, " ")
		if len(segs) != 2 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenStr := segs[1]
		userClaims := &web.UserClaims{}
		// 2. 解析token
		token, err := jwt.ParseWithClaims(tokenStr, userClaims, func(token *jwt.Token) (interface{}, error) {
			return []byte("SUwcr3HfInY49a4XVQ03lV4u1AgcQkynTkf5dPbEAknqr8K7zh5WFNLLPgpUocHi"), nil
		})
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if token == nil || !token.Valid || userClaims.Uid == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Set("userClaims", userClaims)
	}
}
