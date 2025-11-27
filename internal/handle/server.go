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
	r.Use(gin.Recovery(), gin.Logger(), RequestTimeout(), CorsMiddleware())
	// 注册路由
	RegisterRouter(r)
	// 启动HTTP服务器
	s.startHTTPServer(r)
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
