package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

type TestJob1 struct {
}

func (t TestJob1) Run() {
	fmt.Println("testob1.....")
}

type TestJob2 struct {
}

func (t TestJob2) Run() {
	fmt.Println("testob2.....")
}

func main() {
	c := cron.New()
	spec:="*/2 * * * * ?"
	c.AddFunc(spec, func() {
		fmt.Println(time.Now().Format("2006-01-02:15:04:05"))
	})
	c.AddJob(spec,TestJob1{})
	c.AddJob(spec,TestJob2{})

	c.Run()
	select {}
}
