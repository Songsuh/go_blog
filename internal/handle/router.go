package handle

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.Engine) {
	// 注册路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
}

//func setGroup(r *gin.RouterGroup) {
//	noAuth := r.Group("/no-auth")
//}

// registerNoAuthRouter 注册无需认证的路由
func registerNoAuthRouter(r *gin.Engine) {
	noAuth := r.Group("/no-auth")
	{
		noAuth.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "no-auth ping",
			})
		})
	}
}
