package gredis

import (
	"fmt"
	"os"
	"utils/db/gredis/gredis_config"
	"utils/third_party/go-redis"
)

const (
	StandaloneMode string = "StandaloneMode" // 单机redis
	ClusterMode    string = "ClusterMode"    // 集群redis
)

var (
	MainRdsConn *RdsCon // 主服务器对应的Redis
)

func Setup(mode string, staSetting *gredis_config.Redis, cluSetting *gredis_config.RedisClu) {
	MainRdsConn = NewRdsCon(mode, staSetting, cluSetting) // 集群
	if MainRdsConn.cli == nil {
		fmt.Print("connect Redis Failed")
		os.Exit(3)
	}
}

func NewRdsCon(mode string, staSetting *gredis_config.Redis, cluSetting *gredis_config.RedisClu) *RdsCon {
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

func newStandaloneClient(rdsSetting *gredis_config.Redis) redis.Cmdable { //nolint:ireturn
	client := redis.NewClient(&redis.Options{
		Addr:         rdsSetting.Host,
		ReadTimeout:  rdsSetting.ReadTimeout,
		WriteTimeout: rdsSetting.WriteTimeout,
		DialTimeout:  rdsSetting.DialTimeout,
		MinIdleConns: rdsSetting.MaxIdle,
		PoolSize:     rdsSetting.PoolSize,
		DB:           rdsSetting.DB,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return client
}

func newClusterClient(cluRdsSetting *gredis_config.RedisClu) redis.Cmdable {
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
			cluRdsSetting.Slot1Slave,
			cluRdsSetting.Slot2Slave,
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
