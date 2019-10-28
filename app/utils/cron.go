//@Version: V1.0
//@Author: wevsmy
//@License: Apache Licence
//@Contact: wevsmy@gmail.com
//@Site: blog.weii.ink
//@Software: GoLand
//@File: cron.go
//@Time: 2019/10/22 下午4:26

package utils

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

//定时任务

func T() {
	// 给本地时间创建一个 Cron job runner
	c := cron.New(cron.WithSeconds())
	// AddFunc 会向 Cron job runner 添加一个 func ，以按给定的时间表运行  5s执行一次
	spec := "*/5 * * * * ?"
	id, e := c.AddFunc(spec, func() {
		fmt.Println("123456:", time.Now())
	})
	// 解析时间表,是否有问题
	fmt.Println("id:", id, "e：", e)
	// 在当前执行的程序中启动 Cron 调度程序。其实这里的主体是 goroutine + for + select + timer 的调度控制
	c.Start()
	//select {}
	time.Sleep(10 * time.Second)
	id2, e := c.AddFunc(spec, func() {
		fmt.Println("q:", time.Now())
	})
	// 解析时间表,是否有问题
	fmt.Println("id2:", id2, "e：", e)
	for _, entry := range c.Entries() {
		fmt.Println(entry)
		fmt.Println(entry.ID, entry.Job, entry.WrappedJob, entry.Prev, entry.Next)
	}
	fmt.Println("123")
	c.Remove(id)
	time.Sleep(10 * time.Second)
	fmt.Println("==========")
	for _, entry := range c.Entries() {
		fmt.Println(entry)
		fmt.Println(entry.ID, entry.Job, entry.WrappedJob, entry.Prev, entry.Next)
	}
}
