package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		//请求前
		path := c.Request.URL.Path
		method := c.Request.Method
		log.Printf("[INFO]请求开始：%s %s", method, path)

		//请求处理
		c.Next()

		//请求后
		latecy := time.Since(t)
		status := c.Writer.Status()
		log.Printf("[INFO]请求结束:%s %s 状态码=%d 用时=%v ", method, path, status, latecy)
	}
}
