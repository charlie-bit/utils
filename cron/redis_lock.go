package cron

import (
	"errors"
	"log"
	"time"
	"github.com/charlie-bit/utils/third_party/go-redis"
)

const (
	DefaultTTL = 1 * time.Minute
)

type RedisLock struct {
	client *redis.Client
	ttl    time.Duration
	key    string
}

// Lock get a distributed lock
func Lock(client *redis.Client, key string, ttl time.Duration) (*RedisLock, error) {
	if client == nil {
		return nil, errors.New("invalid parameter client")
	}
	if key == "" {
		return nil, errors.New("invalid parameter key")
	}
	if ttl <= 0 {
		ttl = DefaultTTL
	}

	rl := &RedisLock{
		client: client,
		ttl:    ttl,
		key:    key,
	}

	ok, err := rl.client.SetNX(rl.key, time.Now().Unix(), rl.ttl).Result()
	if !ok {
		return nil, errors.New("get lock failed")
	}
	if err != nil {
		log.Printf("redis err:%v", err.Error())
		return nil, err
	}

	return rl, nil
}

// if cron business is successful, I think it would be better for redis key expire to automatically unlock.
func (r *RedisLock) unLock() {
	if r == nil {
		return
	}
	if err := r.client.Del(r.key).Err(); err != nil {
		log.Printf("unlock failed  {key: %v, err: %v}", r.key, err.Error())
	}
}
