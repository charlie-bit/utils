package gredis

import (
	"errors"
	"time"
	"utils/third_party/go-redis"
)

const RdsNil = "redis: nil"

type RdsCon struct {
	cli redis.Cmdable
}

func (self *RdsCon) IsConnected() bool {
	return self.cli != nil
}

// 为空返回0，否则返回1
// 禁止先exist再get，直接get来判空
func (self *RdsCon) Exist(key string) (bool, error) {
	r, err := self.cli.Exists(key).Result()
	if err != nil {
		return false, err
	}
	if r == 1 {
		return true, nil
	}
	return false, nil
}

// 删除成功返回true，不存在时返回false
// 删除大key会阻塞，用unlink，严禁用Del
func (self *RdsCon) Del(key string) (bool, error) {
	r, err := self.cli.Del(key).Result()
	if r == 0 {
		return false, err
	}
	return true, err
}

func (self *RdsCon) Unlink(keys ...string) (int64, error) {
	return self.cli.Unlink(keys...).Result()
}

// 持续多久后消失,尽量不用,设置时就指定
func (self *RdsCon) Expire(key string, expiration time.Duration) (bool, error) {
	return self.cli.Expire(key, expiration).Result()
}

// 持续到某个时间戳,尽量不用,设置时就指定
func (self *RdsCon) ExpireAt(key string, ts int64) (bool, error) {
	return self.cli.ExpireAt(key, time.Unix(ts, 0)).Result()
}

func (self *RdsCon) Rename(key, newkey string) (string, error) {
	return self.cli.Rename(key, newkey).Result()
}

// 获取TTL结束时间
func (self *RdsCon) TTL(key string) (time.Duration, error) {
	return self.cli.TTL(key).Result()
}

/*
	string
*/
// 不存在时返回空
func (self *RdsCon) Get(key string) (string, error) {
	rs, e := self.cli.Get(key).Result()
	if e != nil && e.Error() == RdsNil {
		return "", nil
	}
	return rs, e
}

func (self *RdsCon) GetSet(key string, value interface{}) (string, error) {
	return self.cli.GetSet(key, value).Result()
}

// 第二个参数代表value，第三个代表持续时间（没有默认0）
// value只能存String,uint,int8,int64,int32,float64格式,时间只能为time.Duration类型
// eg: gredis.MainRdsConn.set("1","2")| gredis.MainRdsConn.set("1","2",time.Second(10))
func (self *RdsCon) Set(key string, value ...interface{}) (string, error) {
	switch value[0].(type) {
	case int, int8, int32, int64, uint, uint8, uint32, uint64, float32, float64, string:
		if len(value) == 2 {
			switch value[1].(type) {
			case time.Duration:
				return self.cli.Set(key, value[0], value[1].(time.Duration)).Result()
			}
		} else {
			return self.cli.Set(key, value[0], time.Duration(0)).Result()
		}
	}
	return "", errors.New("wrong type")
}

// 设置成功（原来没有）为true，否则为false
func (self *RdsCon) SetNx(key string, value interface{}, expiration time.Duration) (bool, error) {
	return self.cli.SetNX(key, value, expiration).Result()
}

func (self *RdsCon) SetXx(key string, value interface{}, expiration time.Duration) (bool, error) {
	return self.cli.SetXX(key, value, expiration).Result()
}

// 增加并返回增加后的值
func (self *RdsCon) IncrBy(key string, value int64) (int64, error) {
	return self.cli.IncrBy(key, value).Result()
}

/*
	HASH
*/

func (self *RdsCon) HGet(key, field string) (string, error) {
	r, e := self.cli.HGet(key, field).Result()
	if e != nil && e.Error() == RdsNil {
		return "", nil
	}
	return r, e
}

func (self *RdsCon) HMGet(key string, fields ...string) ([]interface{}, error) {
	if len(fields) == 0 {
		return []interface{}{}, nil
	}
	return self.cli.HMGet(key, fields...).Result()
}

func (self *RdsCon) HGetAll(key string) (map[string]string, error) {
	return self.cli.HGetAll(key).Result()
}

func (self *RdsCon) HExists(key, field string) (bool, error) {
	return self.cli.HExists(key, field).Result()
}

// redis里迭代。查找特定field时代替hgetall
func (self *RdsCon) HScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return self.cli.HScan(key, cursor, match, count).Result()
}

func (self *RdsCon) HSet(key, field string, value string) (bool, error) {
	r, e := self.cli.HSet(key, field, value).Result()
	if r > 0 {
		return true, e
	} else {
		return false, e
	}
}

// 可设置多组value,返回设置成功了的数量
// 没有格式校验，少用
// HSet accepts values in following formats:
//   - HSet("myhash", "key1", "value1", "key2", "value2")
//   - HSet("myhash", []string{"key1", "value1", "key2", "value2"})
//   - HSet("myhash", map[string]interface{}{"key1": "value1", "key2": "value2"})
func (self *RdsCon) HSetInterface(key string, values ...interface{}) (int64, error) {
	return self.cli.HSet(key, values...).Result()
}

func (self *RdsCon) HDel(key string, field ...string) (int64, error) {
	return self.cli.HDel(key, field...).Result()
}

func (self *RdsCon) HSetNx(key, field string, value interface{}) (bool, error) {
	return self.cli.HSetNX(key, field, value).Result()
}

// 返回设置成功了的数量
func (self *RdsCon) HMSet(key string, fields map[string]interface{}) (int64, error) {
	if len(fields) == 0 {
		return 0, nil
	}
	return self.cli.HSet(key, fields).Result()
}

func (self *RdsCon) HLen(key string) (int64, error) {
	return self.cli.HLen(key).Result()
}

// 增加
func (self *RdsCon) HIncrBy(key, field string, value int64) (int64, error) {
	return self.cli.HIncrBy(key, field, value).Result()
}

func (self *RdsCon) HIncrByFloat(key, field string, value float64) (float64, error) {
	return self.cli.HIncrByFloat(key, field, value).Result()
}

/*
	ZSET
*/
// 不需要判断存不存在
func (self *RdsCon) ZCard(key string) (int64, error) {
	return self.cli.ZCard(key).Result()
}

func (self *RdsCon) ZAdd(key string, members ...*redis.Z) (int64, error) {
	return self.cli.ZAdd(key, members...).Result()
}

func (self *RdsCon) ZRem(key string, members ...interface{}) (int64, error) {
	return self.cli.ZRem(key, members...).Result()
}

func (self *RdsCon) ZCount(key, min, max string) (int64, error) {
	return self.cli.ZCount(key, min, max).Result()
}

func (self *RdsCon) ZRank(key string, member string) (int64, error) {
	r, e := self.cli.ZRank(key, member).Result()
	if e != nil && e.Error() == RdsNil {
		return -1, nil
	}
	return r, e
}

func (self *RdsCon) ZScore(key string, member string) (float64, error) {
	r, e := self.cli.ZScore(key, member).Result()
	if e != nil && e.Error() == RdsNil {
		return 0, nil
	}
	return r, e
}

func (self *RdsCon) ZScoreWithError(key string, member string) (float64, error) {
	return self.cli.ZScore(key, member).Result()
}

func (self *RdsCon) ZIncr(key string, member string, incr float64) (float64, error) {
	return self.cli.ZIncr(key, &redis.Z{Score: incr, Member: member}).Result()
}

func (self *RdsCon) ZRevRank(key string, member string) (int64, error) {
	r, e := self.cli.ZRevRank(key, member).Result()
	if e != nil && e.Error() == RdsNil {
		return -1, nil
	}
	return r, e
}

// 获取排名和积分 isRev 为true 从大到小，为false 从小到大
func (self *RdsCon) ZRankWithScore(key string, member string, isRev bool) (int64, float64, error) {
	var rank int64
	var err error
	if isRev {
		rank, err = self.cli.ZRevRank(key, member).Result()
	} else {
		rank, err = self.cli.ZRank(key, member).Result()
	}
	if err != nil {
		if err.Error() == RdsNil {
			return -1, 0, nil
		}
		return 0, 0, err
	}

	score, err := self.cli.ZScore(key, member).Result()
	if err != nil {
		if err.Error() == RdsNil {
			return -1, 0, nil
		}
		return 0, 0, err
	}
	return rank, score, nil
}

// withScoreFlag=true时用[]redis.Z来装，否则用[]string来装
func (self *RdsCon) ZRange(key string, start, stop int64, withScoreFlag bool) (interface{}, error) {
	if withScoreFlag {
		return self.cli.ZRangeWithScores(key, start, stop).Result()
	}
	return self.cli.ZRange(key, start, stop).Result()
}

func (self *RdsCon) ZRevRange(key string, start, stop int64, withScoreFlag bool) (interface{}, error) {
	if withScoreFlag {
		return self.cli.ZRevRangeWithScores(key, start, stop).Result()
	}
	return self.cli.ZRevRange(key, start, stop).Result()
}

// 取全部为(inf,+inf,0,int64_max)
// withScoreFlag=true时用[]redis.Z来装，否则用[]string来装
func (self *RdsCon) ZRangeByScore(key string, minValue, maxValue string, offset, count int64, withScoreFlag bool) (interface{}, error) {
	opt := redis.ZRangeBy{Min: minValue, Max: maxValue, Offset: offset, Count: count}
	if withScoreFlag {
		return self.cli.ZRangeByScoreWithScores(key, &opt).Result()
	}
	return self.cli.ZRangeByScore(key, &opt).Result()
}

func (self *RdsCon) ZRevRangeByScore(key string, minValue, maxValue string, offset, count int64, withScoreFlag bool) (interface{}, error) {
	opt := redis.ZRangeBy{Min: minValue, Max: maxValue, Offset: offset, Count: count}
	if withScoreFlag {
		return self.cli.ZRevRangeByScoreWithScores(key, &opt).Result()
	}
	return self.cli.ZRevRangeByScore(key, &opt).Result()
}

/*
	SET
*/
// 已存在返回false
func (self *RdsCon) SAdd(key string, members ...interface{}) (bool, error) {
	r, e := self.cli.SAdd(key, members...).Result()
	if r == 0 {
		return false, e
	}
	return true, e
}

func (self *RdsCon) SRem(key string, members ...interface{}) (bool, error) {
	r, e := self.cli.SRem(key, members...).Result()
	if r == 0 {
		return false, e
	}
	return true, e
}

func (self *RdsCon) SCard(key string) (int64, error) {
	return self.cli.SCard(key).Result()
}

func (self *RdsCon) SDiff(keys1, key2 string) ([]string, error) {
	return self.cli.SDiff(keys1, key2).Result()
}

func (self *RdsCon) SIsMember(key, member string) (bool, error) {
	return self.cli.SIsMember(key, member).Result()
}

func (self *RdsCon) Smembers(key string) (map[string]struct{}, error) {
	return self.cli.SMembersMap(key).Result()
}

/*
	pipeline
*/
// 严禁pipeline里用hmget等multi命令
func (self *RdsCon) Pipeline() redis.Pipeliner {
	return self.cli.Pipeline()
}

/*
bitmap
*/
func (self *RdsCon) SetBit(key string, offset int64, value int) (int64, error) {
	if offset > 2000000000 { // 原生bitmap最大512MB
		return 0, errors.New("offset too long")
	}
	if value != 0 {
		value = 1
	}
	return self.cli.SetBit(key, offset, value).Result()
}

func (self *RdsCon) GetBit(key string, offset int64) (int64, error) {
	if offset > 2000000000 {
		return 0, errors.New("offset too long")
	}
	return self.cli.GetBit(key, offset).Result()
}

func (self *RdsCon) ZRemRangeByScore(key, min, max string) (int64, error) {
	return self.cli.ZRemRangeByScore(key, min, max).Result()
}
