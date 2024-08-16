package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		//c.String(http.StatusOK, "hello world")
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "ok",
			"data": "hello world",
		})
	})
	fmt.Println("http://localhost:80")
	r.Run(":80")
}
