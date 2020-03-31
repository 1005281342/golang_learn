package main

import (
	"fmt"
	"time"
)

func main() {
	timeDemo()
	timestampDemo()
	timestamp2timeDemo(1585668708)
	timeOpDemo()
	beforeAndAfterDemo()
	tickDemo()
	formatTimeDemo()
	date2timestamp()
}

func timeDemo() {
	// 获取当前时间
	now := time.Now()
	fmt.Printf("current time: %v\n", now)

	// 年
	fmt.Println(now.Year())
	// 月
	fmt.Println(now.Month())
	// 日
	fmt.Println(now.Day())
	// 时
	fmt.Println(now.Hour())
	// 分
	fmt.Println(now.Minute())
	// 秒
	fmt.Println(now.Second())
}

// 时间戳
func timestampDemo() {
	now := time.Now()

	// 时间戳：单位秒
	fmt.Println(now.Unix())

	// 时间戳：单位纳秒
	fmt.Println(now.UnixNano())
}

// 时间戳(秒)转时间
func timestamp2timeDemo(timestamp int64) {
	timeObj := time.Unix(timestamp, 0)
	fmt.Println(timeObj)
	fmt.Println(timeObj.Year())
	fmt.Println(timeObj.Month())
	fmt.Println(timeObj.Day())
	fmt.Println(timeObj.Hour())
	fmt.Println(timeObj.Minute())
	fmt.Println(timeObj.Second())
}

/* time包中
const (
    Nanosecond  Duration = 1
    Microsecond          = 1000 * Nanosecond
    Millisecond          = 1000 * Microsecond
    Second               = 1000 * Millisecond
    Minute               = 60 * Second
    Hour                 = 60 * Minute
)
*/

// 时间操作加减
func timeOpDemo() {
	now := time.Now()
	fmt.Printf("now: %v \n", now)
	now1 := now.Add(time.Hour)
	fmt.Printf("now +1: %v \n", now1)
	a := now1.Sub(now)
	fmt.Printf("a: %v \n", a)
	if a != time.Hour {
		fmt.Println("逻辑异常")
	}
}

// 时间比较
func beforeAndAfterDemo() {
	now := time.Now()
	now1 := now.Add(time.Hour)
	if now1.Before(now) {
		fmt.Println("逻辑错误")
	}
	if now.After(now1) {
		fmt.Println("逻辑错误")
	}
}

// 定时器
func tickDemo() {
	ticker := time.Tick(time.Second)
	cnt := 0
	for i := range ticker {
		if cnt > 2 {
			break
		}
		cnt += 1
		fmt.Println(i)
	}
}

// 时间格式化
func formatTimeDemo() {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05"))
}

// 日期转时间戳
func date2timestamp() {
	fmt.Println(time.Parse("2006-01-02 15:04:05", "2019-03-31 23:58:55"))
}
