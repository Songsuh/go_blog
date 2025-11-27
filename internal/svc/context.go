package svc

import (
	"fmt"
	"github.com/Songsuh/go_blog/internal/global"
	"github.com/Songsuh/go_blog/internal/tools"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var (
	Conf     *global.Config   // 全局配置指针
	onceInit = new(sync.Once) // 确保只初始化一次
	onceGet  = new(sync.Once) // 确保只获取一次配置
	svcCtx   *ServiceContext  // 全局服务上下文
)

type ServiceContext struct {
	Config *global.Config      // 全局配置指针
	Redis  *redis.Client       // Redis客户端
	Mysql  map[string]*gorm.DB // 多数据库连接映射
}

// ReadConfig
/*
	1.读取配置文件
	2.解析配置文件
	3.返回配置
*/
func ReadConfig() *global.Config {
	onceInit.Do(func() {
		// 确保Conf被初始化，避免空指针错误
		if Conf == nil {
			Conf = &global.Config{}
		}
		rootDir, err := tools.GetRootDir()
		if err != nil {
			log.Fatal(err)
		}
		err = godotenv.Load(filepath.Join(rootDir, ".env"))
		if err != nil {
			panic(".env file not found")
		}
		appEnv := os.Getenv("APP_ENV")
		configFile := fmt.Sprintf("%s/etc/config-%s.yaml", rootDir, appEnv)
		v := viper.New()
		v.SetConfigFile(configFile)
		v.SetConfigType("yaml") // 显式设置配置文件类型，增加可靠性
		if err := v.ReadInConfig(); err != nil {
			log.Fatalf("读取配置文件失败: %v", err)
		}
		if err := v.Unmarshal(&Conf); err != nil {
			log.Fatalf("解析配置文件失败: %v", err)
		}
		log.Println("配置文件内容加载成功: ", configFile)
		// 初始化日志系统
		InitLogger(&Conf.Log)
	})
	return Conf
}

func newSvc() *ServiceContext {
	return &ServiceContext{
		Config: ReadConfig(),
		Redis:  CreateRedis(&Conf.Redis),
		Mysql:  CreateMysqls(Conf.Mysql),
	}
}

// GetSvc 获取配置
func GetSvc() *ServiceContext {
	if svcCtx == nil {
		onceGet.Do(func() {
			svcCtx = newSvc()
		})
	}
	return svcCtx
}

// GetDb 获取指定数据库连接
func (s *ServiceContext) GetDb(alias string) *gorm.DB {
	if db, exists := s.Mysql[alias]; exists {
		return db
	}
	log.Printf("警告: 数据库别名 '%s' 不存在", alias)
	return nil
}
