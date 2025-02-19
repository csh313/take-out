package middle

import (
	"github.com/gin-gonic/gin"
	"hmshop/common/enum"
	"hmshop/global"
	pwd "hmshop/utils"
	"net/http"
	"strconv"
)

// 中间件是为了过滤路由而发明的一种机制,也就是http请求来到时先经过中间件,再到具体的处理函数。
// 在中间件中使用协程  主程序会等待协程执行完成，不需要sync
// gin.HandlerFunc其实返回的是一个以c *gin.Context为参数的函数
func AuthMiddleWare() gin.HandlerFunc {

	return func(c *gin.Context) {
		//获取token
		token := c.GetHeader(global.AppConfig.Jwt.Admin.Name)
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "没有传入token",
			})
			c.Abort() //提前结束请求处理
			return
		}
		claims, err := pwd.ParseToken(token, global.AppConfig.Jwt.Admin.Secret)
		if err != nil {
			c.JSON(401, gin.H{"error": "无效的token"})
			c.Abort()
			return
		}

		// 检查 Redis 中的 Token 是否有效
		userID := claims.UserId
		redisToken, err := global.Redis.Get(c, "token:"+token).Result()
		uid, _ := strconv.ParseUint(redisToken, 10, 64)
		if err != nil || uid != userID {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired or invalid"})
			c.Abort()
			return
		}
		//Set用于存储此上下文专用的新键值对
		c.Set(enum.CurrentId, claims.UserId)

		c.Set(enum.CurrentName, claims.GrantScope)

		//验证后处理剩下的程序
		c.Next()
	}
}

func AuthUserMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(global.AppConfig.Jwt.User.Name)
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "没有传入Authorization",
			})
			c.Abort() //提前结束请求处理
			return
		}
		claims, err := pwd.ParseToken(token, global.AppConfig.Jwt.User.Secret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "无效的token",
			})
			c.Abort()
			return
		}
		userID := claims.UserId
		reedisToken, err := global.Redis.Get(c, "token:"+token).Result()
		uid, _ := strconv.ParseUint(reedisToken, 10, 64)
		if err != nil || uid != userID {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token expired or invalid",
			})
			c.Abort()
			return
		}

		c.Set(enum.CurrentUserId, userID)
		c.Set(enum.CurrentUserName, claims.GrantScope)
		//验证后处理剩下的程序
		c.Next()
	}
}
