package chapter4

import (
	"fmt"
	"time"
)

//time包提供了一个数据类型time.Time（作为值来使用）
//以及显示和测量时间和日期的功能函数
//当前时间使用time.Now()获取，或者使用t.Day()，t.Minute()
//等等来获取时间的一部分，甚至可以自定义时间格式化字符串
func TestOutput(){
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%4d\n",t.Day(),t.Month(),t.Year())
	//注意：这里的02d中的0是2位不足使用0来填充

	//测试格式化输出的时间字符串
	fmt.Println(t.Format("02 Jan 2006 15:04"))
	t = time.Now().UTC()
	fmt.Println(t)

	//计算时间
	week := 60*60*24*7*1e9 //必须使用纳秒
	week_from_now := t.Add(time.Duration(week))
	fmt.Println(week_from_now)
	//预定义的时间格式化字符串
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))


}
//Duration类型表示两个连续时刻所相差的纳秒数，类型是int64。Location类型映射某个时区的时间
//UTC表示通用协调世界时间
//
//包中的一个预定义函数func (t Time) Format(layout string) string 可以根据一个格式化
//字符串来将一个时间t转换为相应格式的字符串，可以使用一些预定义的格式 如time.ANSIC或time.RFC822
//一般的格式化设计

//
//如果需要在应用程序在经过一定时间或周期执行某项任务，则可以使用time.After
//或time.Ticker 另外 time.Sleep(Duration d)可以实现对某个进程（实质是goroutine）
//时长为d的暂停
