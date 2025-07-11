package middleware

import (
	"GoBlogProject/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	// "github.com/golang-jwt/jwt/v5"
)

// 认证中间键
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")     //获取请求中Authorization的值
		if !strings.HasPrefix(auth, "Bearer ") { //验证是否存在token的值,注意Bearer后面是有空格的
			// c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供token"}) //输出json数据
			utils.ResponseError(c, http.StatusUnauthorized, "未提供token")
			c.Abort() //终止进程链
			return
		}
		tokenStr := strings.TrimPrefix(auth, "Bearer ") //得到token字符串
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(utils.JwtKey), nil
		})
		if err != nil || token == nil {
			// c.JSON(401, gin.H{"error": "Token 解析失败"})
			utils.ResponseError(c, http.StatusUnauthorized, "Token 解析失败")
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			//类型断言加转换
			c.Set("user_id", int(claims["user_id"].(float64)))
			c.Next()
		} else {
			// c.JSON(http.StatusUnauthorized, gin.H{"error": "无效token"})
			utils.ResponseError(c, http.StatusUnauthorized, "无效token")
			c.Abort()
		}
	}

}
