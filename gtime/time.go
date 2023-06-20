package gtime

import (
	"time"
)

// GetCurrentDateYYMMDDMMSS 获取当前的时间 20060102150405
func GetCurrentDateYYMMDDMMSS() string {
	formatter := MakeFormatter(time.Now())
	formatter.maskSlice = []string{LongYearToken, ZeroMonthToken,
		ZeroDayToken, HourToken, ZeroMinuteToken, ZeroSecondToken}
	return formatter.Convert()
}

// GetCurrentDateYYSlMMSlDD 获取当前的时间 2006/01/02 15:04:05
func GetCurrentDateYYSlMMSlDD() string {
	formatter := MakeFormatter(time.Now())
	formatter.maskSlice = []string{LongYearToken + "/" + ZeroMonthToken + "/" +
		ZeroDayToken + " " + HourToken + ":" + ZeroMinuteToken + ":" + ZeroSecondToken}
	return formatter.Convert()
}

// GetCurrentUnix 获取当前的时间 - Unix时间戳
func GetCurrentUnix() int64 {
	return time.Now().Unix()
}

// GetCurrentMilliUnix 获取当前的时间 - 毫秒级时间戳
func GetCurrentMilliUnix() int64 {
	return time.Now().UnixNano() / 1000000
}

// GetCurrentNanoUnix 获取当前的时间 - 纳秒级时间戳
func GetCurrentNanoUnix() int64 {
	return time.Now().UnixNano()
}
