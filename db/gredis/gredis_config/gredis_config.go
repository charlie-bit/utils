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
