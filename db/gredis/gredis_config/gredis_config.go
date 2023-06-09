package gredis_config

import "time"

type Redis struct {
	Host         string
	Password     string
	MaxIdle      int
	MaxActive    int
	IdleTimeout  time.Duration
	PoolSize     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	DialTimeout  time.Duration
	DB           int
}

type RedisClu struct {
	Slot1        string
	Slot1Start   int
	Slot1End     int
	Slot1Slave   string
	Slot2        string
	Slot2Start   int
	Slot2End     int
	Slot2Slave   string
	MaxRedirects int

	RouteByLatency          bool
	RouteRandomly           bool
	PoolSize                int
	MinIdleConns            int
	PoolTimeout             time.Duration // 等待获取连接的最大时长
	IdleCheckFrequency      time.Duration // 闲置连接检查的周期
	MaxRetries              int
	DialTimeout             time.Duration // DialTimeout 拨超时时间
	ReadTimeout             time.Duration // ReadTimeout 读超时 默认3s
	WriteTimeout            time.Duration // WriteTimeout 读超时 默认3s
	IdleTimeout             time.Duration // IdleTimeout 连接最大空闲时间，默认60s, 超过该时间，连接会被主动关闭
	ReadOnly                bool          // ReadOnly 集群模式 在从属节点上启用读模式
	EnableMetricInterceptor bool          // 是否开启监控，默认开启
	EnableTraceInterceptor  bool          // 是否开启链路，默认开启
}
