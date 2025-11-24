package global

import (
	"time"
)

// Config 应用配置结构体
type Config struct {
	Server ServerConfig     `yaml:"Server"`
	Log    LogConfig        `yaml:"Log"`
	Jwt    JwtConfig        `yaml:"Jwt"`
	Redis  RedisConfig      `yaml:"Redis"`
	Mysql  map[string]Mysql `yaml:"Mysql"`
	Upload UploadConfig     `yaml:"Upload"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port string `yaml:"Port"`
	Name string `yaml:"Name"`
	Mode string `yaml:"Mode"`
}

// LogConfig 日志配置
type LogConfig struct {
	Name     string `yaml:"Name"`
	Encoding string `yaml:"Encoding"`
	Stat     bool   `yaml:"Stat"`
	Mode     string `yaml:"Mode"`
	Level    string `yaml:"Level"`
	Compress bool   `yaml:"Compress"`
	KeepDays int    `yaml:"KeepDays"`
	Rotation string `yaml:"Rotation"`
	Path     string `yaml:"Path"`
}

// JwtConfig JWT配置
type JwtConfig struct {
	Secret string        `yaml:"Secret"`
	Expire time.Duration `yaml:"Expire"`
	Issuer string        `yaml:"Issuer"`
}

// RedisConfig Redis配置
type RedisConfig struct {
	Addr         string        `yaml:"Addr"`
	Username     string        `yaml:"Username"`
	Password     string        `yaml:"Password"`
	Database     int           `yaml:"Database"`
	ClientName   string        `yaml:"ClientName"`
	MaxRetries   int           `yaml:"MaxRetries"`
	ReadTimeout  time.Duration `yaml:"ReadTimeout"`
	WriteTimeout time.Duration `yaml:"WriteTimeout"`
	dialTimeout  time.Duration `yaml:"dialTimeout"`
	PoolSize     int           `yaml:"PoolSize"`
	MinIdleConns int           `yaml:"MinIdleConns"`
	MaxIdleConns int           `yaml:"MaxIdleConns"`
}

// Mysql 数据库配置
type Mysql struct {
	Desc     string `yaml:"desc"`
	Type     string `yaml:"type"`
	Hostname string `yaml:"hostname"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
	Charset  string `yaml:"charset"`
	Prefix   string `yaml:"prefix"`
	Pool     DbPool `yaml:"pool,omitempty"`
}
type DbPool struct {
	MaxIdleConn int           `yaml:"maxIdleConn,omitempty"`
	MaxOpenConn int           `yaml:"maxOpenConn,omitempty"`
	MaxLifetime time.Duration `yaml:"maxLifetime,omitempty"`
}

type MongoDb struct {
	Database        string        `yaml:"database"`
	Uri             string        `yaml:"uri"`
	Username        string        `yaml:"username"`
	Password        string        `yaml:"password"`
	Mechanism       string        `yaml:"mechanism"`
	AuthSource      string        `yaml:"authSource"`
	ConnectTimeout  time.Duration `yaml:"connectTimeout"`
	MaxPoolSize     uint64        `yaml:"maxPoolSize"`
	MinPoolSize     uint64        `yaml:"minPoolSize"`
	Timeout         time.Duration `yaml:"timeout"`
	MaxConnIdleTime time.Duration `yaml:"maxConnIdleTime"`
}

// UploadConfig 上传配置
type UploadConfig struct {
	Path string `yaml:"Path"`
}
