package time

import (
	"fmt"
	"time"
)

// 获取当前的时间 20060102150405
func GetCurrentDateYYMMDDMMSS() string {
	formatter := MakeFormatter(time.Now())
	formatter.maskSlice = []string{LongYearToken, ZeroMonthToken,
		ZeroDayToken, HourToken, ZeroMinuteToken, ZeroSecondToken}
	return formatter.Convert()
}

// 获取当前的时间 2006/01/02 15:04:05
func GetCurrentDateYYSlMMSlDD() string {
	formatter := MakeFormatter(time.Now())
	formatter.maskSlice = []string{LongYearToken + "/" + ZeroMonthToken + "/" +
		ZeroDayToken + " " + HourToken + ":" + ZeroMinuteToken + ":" + ZeroSecondToken}
	fmt.Println(formatter.GetMask())
	return formatter.Convert()
}

// 获取当前的时间 - Unix时间戳
func GetCurrentUnix() int64 {
	return time.Now().Unix()
}

// 获取当前的时间 - 毫秒级时间戳
func GetCurrentMilliUnix() int64 {
	return time.Now().UnixNano() / 1000000
}

// 获取当前的时间 - 纳秒级时间戳
func GetCurrentNanoUnix() int64 {
	return time.Now().UnixNano()
}
