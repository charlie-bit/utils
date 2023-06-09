package gmysql

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"utils/db/common"
)

var emptyCtx = context.Background()

type MysqlClient struct {
	*gorm.DB
}

// Option 可选项
type Option func(c *common.Config)

// NewMysqlClient ...
func NewMysqlClient(config *common.Config, options ...Option) (*MysqlClient, error) {
	mysqlClient := MysqlClient{}

	gormCfg := gorm.Config{}
	// 不开启 raw debug 时, 关闭 gorm 原生日志
	if !config.RawDebug {
		gormCfg.Logger = logger.Discard
	}

	db, err := gorm.Open(mysql.Open(config.DSN), &gormCfg)
	if err != nil {
		return nil, err
	}

	if config.RawDebug {
		db = db.Debug()
	}

	gormDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 设置默认连接配置
	gormDB.SetMaxIdleConns(config.MaxIdleConns)
	gormDB.SetMaxOpenConns(config.MaxOpenConns)

	if config.ConnMaxLifetime != 0 {
		gormDB.SetConnMaxLifetime(config.ConnMaxLifetime)
	}

	mysqlClient.DB = db
	return &mysqlClient, nil
}

// WithContext ...
func (m *MysqlClient) WithContext(ctx context.Context) *MysqlClient {
	m.Statement.Context = ctx
	return m
}
