package server

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.Engine) {
	// 注册路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
	registerNoAuthRouter(r) // 注册无需认证的路由
	registerApiRouter(r)    // 注册API路由
	registerAdminRouter(r)  // 注册管理员路由
}

// registerNoAuthRouter 注册无需认证的路由
func registerNoAuthRouter(r *gin.Engine) {
	r.POST("api/v1/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "no-auth login",
		})
	})
}

// registerApiRouter 注册API路由
func registerApiRouter(r *gin.Engine) {
	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "api-v1 ping",
			})
		})
		apiV1.POST("/info", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "api-v1 info",
			})
		})
	}
}

// registerAdminRouter 注册管理员路由
func registerAdminRouter(r *gin.Engine) {
	admin := r.Group("/admin")
	{
		admin.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "admin ping",
			})
		})
	}
}
