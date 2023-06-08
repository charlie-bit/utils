package grandom

import (
	"math/rand"
	"time"
)

type RandomInter interface {
	RandomInt(n int) int
	RandomInt32(n int32) int32
	RandomInt64(n int64) int64
	GetRandomResult(begin, end int, num int) []int
}

var RandomUtil = _randomUtil{}

type _randomUtil struct {
}

func (_randomUtil) RandomInt(n int) int {
	rand.Seed(time.Now().UnixNano())
	if n <= 0 {
		return 0
	}
	return rand.Intn(n)
}

func (_randomUtil) RandomInt32(n int32) int32 {
	rand.Seed(time.Now().UnixNano())
	if n <= 0 {
		return 0
	}
	return rand.Int31n(n)
}

func (_randomUtil) RandomInt64(n int64) int64 {
	rand.Seed(time.Now().UnixNano())
	if n <= 0 {
		return 0
	}
	return rand.Int63n(n)
}

// [begin,end)
// 范围内随机num个结果，采用相同种子
func (_randomUtil) GetRandomResult(begin, end int, num int) []int {
	if begin > end {
		return nil
	}
	itv := end - begin
	rand.Seed(time.Now().UnixNano())
	var result []int
	for i := 0; i < num; i++ {
		result = append(result, begin+rand.Intn(itv))
	}

	return result
}
