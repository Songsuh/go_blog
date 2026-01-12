package svc

import (
	"context"
	"fmt"
	"go_blog/internal/global"
	"log/slog"
	"os"
	"path/filepath"

	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger 全局日志实例
var Logger *slog.Logger

// InitLogger 初始化日志系统
func InitLogger(config *global.LogConfig) {
	var handler slog.Handler

	// 根据配置选择日志输出方式
	switch config.Driver {
	case "file":
		handler = createFileHandler(config)
	case "console":
		handler = createConsoleHandler(config)
	default:
		handler = createConsoleHandler(config) // 默认使用控制台输出
	}

	// 创建日志实例
	Logger = slog.New(handler)

	// 记录日志系统初始化信息
	Logger.Info("日志系统初始化完成",
		"driver", config.Driver,
		"level", config.Level,
		"path", config.Path)
}

// createFileHandler 创建文件日志处理器
func createFileHandler(config *global.LogConfig) slog.Handler {
	// 确保日志目录存在
	if err := os.MkdirAll(config.Path, 0755); err != nil {
		panic(fmt.Sprintf("创建日志目录失败: %v", err))
	}

	// 配置lumberjack日志轮转
	logFile := filepath.Join(config.Path, config.ServerName+".log")
	lj := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    20,              // 单个日志文件最大大小(MB)
		MaxBackups: 10,              // 最大备份文件数
		MaxAge:     config.KeepDays, // 最大保留天数
		Compress:   config.Compress, // 是否压缩备份文件
		LocalTime:  true,            // 使用本地时间
	}

	// 创建JSON处理器
	opts := &slog.HandlerOptions{
		Level: getLogLevel(config.Level),
	}

	return slog.NewJSONHandler(lj, opts)
}

// createConsoleHandler 创建控制台日志处理器
func createConsoleHandler(config *global.LogConfig) slog.Handler {
	opts := &slog.HandlerOptions{
		Level: getLogLevel(config.Level),
	}

	return slog.NewTextHandler(os.Stdout, opts)
}

// getLogLevel 将字符串日志级别转换为slog.Level
func getLogLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo // 默认info级别
	}
}

// LogContext 日志上下文，包含请求ID等信息
type LogContext struct {
	RequestID string
	UserID    string
	IP        string
	Path      string
	Method    string
}

// WithContext 为日志添加上下文信息
func WithContext(ctx context.Context, logCtx LogContext) *slog.Logger {
	if Logger == nil {
		// 如果日志系统未初始化，使用默认日志
		return slog.Default()
	}

	//attrs := []slog.Attr{
	//	slog.String("request_id", logCtx.RequestID),
	//	slog.String("user_id", logCtx.UserID),
	//	slog.String("ip", logCtx.IP),
	//	slog.String("path", logCtx.Path),
	//	slog.String("method", logCtx.Method),
	//	slog.Time("timestamp", time.Now()),
	//}
	//
	//return Logger.With(attrs...)
	return nil
}

func Debug(msg string, args ...any) {
	if Logger != nil {
		Logger.Debug(msg, args...)
	}
}

func Info(msg string, args ...any) {
	if Logger != nil {
		Logger.Info(msg, args...)
	}
}

func Warn(msg string, args ...any) {
	if Logger != nil {
		Logger.Warn(msg, args...)
	}
}

func Error(msg string, args ...any) {
	if Logger != nil {
		Logger.Error(msg, args...)
	}
}

func Errorf(format string, args ...any) {
	if Logger != nil {
		Logger.Error(fmt.Sprintf(format, args...))
	}
}
