package cron

import (
	"time"
)

type HelloWorld struct {
	_Jobs
	expiration time.Duration
	key        string
}

func (u *HelloWorld) SetLockExpire(duration time.Duration) {
	u.expiration = duration
}

func (u *HelloWorld) Run() {
	// get a distributed lock, if the task is successful, no need to unlock, it would be
	// better for redis key expire to automatically unlock, if the task is failed, you can manually unlock.
	l, err := Lock(nil, u.key, u.expiration)
	if err != nil {
		return
	}
	defer l.unLock()
}

func (u *HelloWorld) Unlock() error {
	//if err := db.RedisClient.Del(context.Background(), u.key).Err(); err != nil {
	//	return err
	//}
	return nil
}
