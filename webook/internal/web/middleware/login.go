package middleware

import (
	"encoding/gob"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type LoginMiddlewareBuilder struct {
	paths []string
}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}

func (l *LoginMiddlewareBuilder) IgnorePaths(paths ...string) *LoginMiddlewareBuilder {
	for _, path := range paths {
		l.paths = append(l.paths, path)
	}
	return l
}

func (l *LoginMiddlewareBuilder) Build() gin.HandlerFunc {
	gob.Register(time.Time{})
	return func(ctx *gin.Context) {
		for _, path := range l.paths {
			if ctx.Request.URL.Path == path {
				return
			}
		}
		sess := sessions.Default(ctx)
		id := sess.Get("userid")
		// 判断是否登录
		if id == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		updateTime := sess.Get("update_time")
		now := time.Now()
		//刚登陆没刷新
		if updateTime == nil {
			fmt.Println("第一次登陆，刷新")
			sess.Set("update_time", now)
			sess.Save()
			return
		}

		updateTimeVal, _ := updateTime.(time.Time)

		// 判断是否超时
		if now.Sub(updateTimeVal) > time.Second*10 {
			fmt.Println("超时，刷新")
			sess.Set("update_time", now)
			sess.Save()
			return
		}
	}
}
