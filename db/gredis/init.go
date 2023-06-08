package gredis

import (
	"fmt"
	"os"
	"time"
	"utils/third_party/go-redis"
)

// 定义配置结构
const (
	StandaloneMode string = "StandaloneMode" // 单机redis
	ClusterMode    string = "ClusterMode"    // 集群redis
)

var (
	MainRdsConn *RdsCon // 主服务器对应的Redis
	TeamRdsConn *RdsCon // 公会服务器对应的Redis
)

func Setup() {
	MainRdsConn = NewRdsCon(setting.ServerSetting.MainRedisMode, setting.OldRedisSetting, setting.MainRedisCluSetting) // 集群
	if MainRdsConn.cli == nil {
		fmt.Print("connect Redis Failed")
		os.Exit(3)
	}
}

func NewRdsCon(mode string, staSetting *setting.Redis, cluSetting *setting.RedisClu) *RdsCon {
	if mode == StandaloneMode {
		return &RdsCon{
			cli: newStandaloneClient(staSetting),
		}
	} else {
		return &RdsCon{
			cli: newClusterClient(cluSetting),
		}
	}
}

func newStandaloneClient(rdsSetting *setting.Redis) redis.Cmdable { //nolint:ireturn
	client := redis.NewClient(&redis.Options{
		Addr:         rdsSetting.Host,
		ReadTimeout:  rdsSetting.ReadTimeout,
		WriteTimeout: rdsSetting.WriteTimeout,
		DialTimeout:  rdsSetting.DialTimeout,
		MinIdleConns: rdsSetting.MaxIdle,
		PoolSize:     rdsSetting.PoolSize,
		DB:           rdsSetting.DB,
	})

	// ping5次，都不通认为失败
	for i := 0; i < 5; i++ {
		_, err := client.Ping().Result()
		if err == nil {
			return client
		}
		time.Sleep(100 * time.Microsecond)
	}
	return nil
}

func newClusterClient(cluRdsSetting *setting.RedisClu) redis.Cmdable {
	c := redis.NewClusterClient(&redis.ClusterOptions{
		MaxRedirects: cluRdsSetting.MaxRedirects,

		ReadOnly:       cluRdsSetting.ReadOnly,
		RouteByLatency: cluRdsSetting.RouteByLatency,
		RouteRandomly:  cluRdsSetting.RouteRandomly,

		ReadTimeout:  cluRdsSetting.ReadTimeout,
		WriteTimeout: cluRdsSetting.WriteTimeout,
		DialTimeout:  cluRdsSetting.DialTimeout,

		PoolSize:           cluRdsSetting.PoolSize,
		MinIdleConns:       cluRdsSetting.MinIdleConns,
		PoolTimeout:        cluRdsSetting.PoolTimeout,
		IdleTimeout:        cluRdsSetting.IdleTimeout,
		IdleCheckFrequency: cluRdsSetting.IdleCheckFrequency,

		MaxRetries: cluRdsSetting.MaxRetries,
		Addrs: []string{
			cluRdsSetting.Slot1,
			cluRdsSetting.Slot2,
			cluRdsSetting.Slot3,
			cluRdsSetting.Slot1Slave,
			cluRdsSetting.Slot2Slave,
			cluRdsSetting.Slot3Slave,
		},
	})

	err := c.ForEachNode(func(node *redis.Client) error {
		return node.Ping().Err()
	})
	if err != nil {
		panic(err)
	}

	return c
}
