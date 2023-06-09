package common

import (
	"time"
)

// Config options
type Config struct {
	Name                       string
	Dialect                    string        // 选择数据库种类，默认mysql,postgres
	DSN                        string        // DSN地址: username:password@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=True&loc=Local&timeout=1s&readTimeout=3s&writeTimeout=3s
	Debug                      bool          // 是否开启调试，默认不开启，开启后，可以看到每次请求，配置名、地址、耗时、请求数据、响应数据, 标准输出, 生产请使用 access log
	RawDebug                   bool          // 是否开启原生调试开关，默认不开启
	MaxIdleConns               int           // 最大空闲连接数，默认10
	MaxOpenConns               int           // 最大活动连接数，默认100
	ConnMaxLifetime            time.Duration // 连接的最大存活时间，默认300s
	OnFail                     string        // 创建连接的错误级别，=panic时，如果创建失败，立即panic，默认连接不上panic
	SlowLogThreshold           time.Duration // 慢日志阈值，默认500ms
	EnableMetricInterceptor    bool          // 是否开启监控，默认开启
	EnableTraceInterceptor     bool          // 是否开启链路追踪，默认开启
	EnableDetailSQL            bool          // 是否打印包含参数的完整sql语句，select * from aid = ?;
	EnableAccessInterceptor    bool          // 是否开启，记录请求数据
	EnableAccessInterceptorReq bool          // 是否开启记录请求参数
	EnableAccessInterceptorRes bool          // 是否开启记录响应参数
	EnableRecordNotFoundLog    bool          // ErrRecordNotFound 错误时是否打印 warn 日志, 默认开启
	// Deprecated: not affect anything
	EnableSkyWalking bool // 是否额外开启 skywalking, 默认关闭
}

type MongoConfig struct {
	DBUri                string   `yaml:"dbUri"`
	DBAddress            []string `yaml:"dbAddress"`
	DBDirect             bool     `yaml:"dbDirect"`
	DBTimeout            int      `yaml:"dbTimeout"`
	DBDatabase           string   `yaml:"dbDatabase"`
	DBSource             string   `yaml:"dbSource"`
	DBUserName           string   `yaml:"dbUserName"`
	DBPassword           string   `yaml:"dbPassword"`
	DBMaxPoolSize        int      `yaml:"dbMaxPoolSize"`
	DBRetainChatRecords  int      `yaml:"dbRetainChatRecords"`
	ChatRecordsClearTime string   `yaml:"chatRecordsClearTime"`
}
