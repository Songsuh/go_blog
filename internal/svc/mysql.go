package svc

import (
	"fmt"
	"github.com/Songsuh/go_blog/internal/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

// CreateMysqls 创建Mysql连接
func CreateMysqls(mysqlConfigs map[string]global.Mysql) map[string]*gorm.DB {
	dbs := make(map[string]*gorm.DB)

	for alias, config := range mysqlConfigs {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
			config.Username,
			config.Password,
			config.Hostname,
			config.Port,
			config.Database,
			config.Charset,
		)

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags), // 输出到控制台
				logger.Config{
					SlowThreshold:             time.Second, // 慢查询阈值（超过1秒视为慢查询）
					LogLevel:                  logger.Info, // 日志级别：Info（打印所有SQL）
					IgnoreRecordNotFoundError: true,        // 忽略记录未找到的错误
					Colorful:                  true,        // 启用彩色输出
				},
			),
			PrepareStmt: true,
			NowFunc: func() time.Time {
				return time.Now().Local()
			},
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   config.Prefix,
				SingularTable: true,
				NoLowerCase:   false,
			},
		})

		if err != nil {
			log.Fatalf("mysql connect error for %s: %v", alias, err)
		}

		// 设置连接池参数
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("mysql init error for %s: %v", alias, err)
		}

		if config.Pool.MaxIdleConn > 0 {
			sqlDB.SetMaxIdleConns(config.Pool.MaxIdleConn)
		}
		if config.Pool.MaxOpenConn > 0 {
			sqlDB.SetMaxOpenConns(config.Pool.MaxOpenConn)
		}
		if config.Pool.MaxLifetime > 0 {
			sqlDB.SetConnMaxLifetime(config.Pool.MaxLifetime)
		}

		// 测试连接
		if err := sqlDB.Ping(); err != nil {
			log.Fatalf("mysql ping error for %s: %v", alias, err)
		}

		dbs[alias] = db
		log.Printf("MySQL连接 %s 创建成功: %s@%s:%d/%s",
			alias, config.Username, config.Hostname, config.Port, config.Database)
	}

	return dbs

}
