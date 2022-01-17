package main

import (
	"fmt"
	"github.com/robfig/cron"
	"math/rand"
	"sync"
	"time"
)

// LocalLimitMap 本地存储map
var LocalLimitMap = sync.Map{}

// LocalLimit 本地限流结构体
type LocalLimit struct {
	Count      int64 // 请求次数
	ExpireTime int64 // 过期时间
}

// Run 定时清除过期数据
func (imp LocalLimit) Run() {
	var localLimit LocalLimit
	LocalLimitMap.Range(func(key, value interface{}) bool {
		localLimit = value.(LocalLimit)
		if time.Now().UnixNano()/1e6 > localLimit.ExpireTime {
			LocalLimitMap.Delete(key)
		}
		return false
	})
	//fmt.Println("总共花费时间(ms)==", time.Now().Sub(now).Nanoseconds())
}

func LocalLimitHandle(cmd string, qps int64) bool {
	key := fmt.Sprintf("%s_%d", cmd, time.Now().Unix())
	value, ok := LocalLimitMap.Load(key)
	if ok {
		var localLimit LocalLimit
		localLimit = value.(LocalLimit)
		if localLimit.Count < qps {
			localLimit.Count++
			LocalLimitMap.Store(key, localLimit)
			fmt.Println("localLimit==", localLimit)
			return true
		}

	} else {
		localLimit := LocalLimit{
			Count:      1,
			ExpireTime: time.Now().Unix() + 2,
		}
		LocalLimitMap.Store(key, localLimit)

		return true

	}
	//
	return false
}

func main() {
	go request()
	c := cron.New()
	spec := "*/2 * * * * ?"
	c.AddJob(spec, LocalLimit{})

	c.Run()
	select {}
}

func request() {
	time.Sleep(time.Second * 1)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 400; i++ {
		cmd := fmt.Sprintf("%s_%d", "cmd", rand.Intn(3))
		ok := LocalLimitHandle(cmd, 10)
		if ok {
			fmt.Printf("access  cmd=%s\n", cmd)
		} else {
			fmt.Printf("refuse  cmd=%s\n", cmd)
		}
		time.Sleep(time.Millisecond*20)
	}
}
