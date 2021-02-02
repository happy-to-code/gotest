package main

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
)

type TestJob struct {
}

func (job TestJob) Run() {
	fmt.Println("testJob1...")
}

type Test2Job struct {
}

func (job2 Test2Job) Run() {
	fmt.Println("testJob2...")
}

// 启动多个任务
func main() {
	i := 0
	c := cron.New()

	// AddFunc
	spec := "*/2 * * * * ?"
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	})

	// AddJob方法
	c.AddJob(spec, TestJob{})
	c.AddJob(spec, Test2Job{})

	// 启动计划任务
	c.Start()

	// 关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer c.Stop()

	select {}
}
