package handle

import (
	"context"
	"errors"
	"github.com/Songsuh/go_blog/internal/svc"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type Server struct {
	Mode     string
	Port     string
	bootTime time.Time // 启动时间
	ctx      context.Context
	cancel   context.CancelFunc
}

func NewServer() *Server {
	return &Server{
		Mode: svc.GetSvc().Config.Server.Mode,
		Port: svc.GetSvc().Config.Server.Port,
	}
}

// Run 启动服务器
func (s *Server) Run(ctx context.Context) {
	s.ctx, s.cancel = context.WithCancel(ctx)
	s.bootTime = time.Now()
	s.boot()
}

// Stop 停止服务器
func (s *Server) Stop(ctx context.Context) error {
	if s.cancel != nil {
		s.cancel()
	}
	return nil
}

// boot 启动gin http服务器
func (s *Server) boot() {
	// 启动服务器
	gin.SetMode(s.Mode)
	// 创建gin实例
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger(), s.requestTimeout(), s.corsMiddleware())
	// 注册路由
	RegisterRouter(r)
	// 注册健康检查路由
	s.registerHealthCheck(r)
	// 启动HTTP服务器
	s.startHTTPServer(r)
}

// requestTimeout 请求超时中间件
func (s *Server) requestTimeout() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// corsMiddleware 跨域中间件
func (s *Server) corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// registerHealthCheck 注册健康检查路由
func (s *Server) registerHealthCheck(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":    "healthy",
			"timestamp": time.Now().Unix(),
			"uptime":    time.Since(s.bootTime).String(),
		})
	})

	r.GET("/ready", func(c *gin.Context) {
		// 这里可以添加数据库连接检查等就绪检查
		c.JSON(200, gin.H{
			"status": "ready",
		})
	})
}

// startHTTPServer 启动HTTP服务器
func (s *Server) startHTTPServer(r *gin.Engine) {
	server := &http.Server{
		Addr:         s.Port,
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// 在goroutine中启动服务器
	go func() {
		log.Printf("Server starting on %s in %s mode", s.Port, s.Mode)

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// 等待关闭信号
	s.waitForShutdown(server)
}

// waitForShutdown 等待优雅关闭
func (s *Server) waitForShutdown(server *http.Server) {
	<-s.ctx.Done()

	log.Println("Server is shutting down...")

	// 创建关闭上下文，最多等待30秒
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 优雅关闭服务器
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
