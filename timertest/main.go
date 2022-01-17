package main

import (
	"fmt"
	"time"
)

// timer 时间到了，只执行一次
func main() {
	test07()
}

func test07() {
	t1:=time.NewTicker(time.Second*2)
	for {
		select {
		case <-t1.C:
			fmt.Println("2秒")
		}
	}
}

func a(ch chan string) {
	time.Sleep(time.Microsecond)
	ch <- "aaa"
}

func write(ch chan string) {
	for {
		select {
		case ch<- "hello":
			//fmt.Println("write hello")
		default:
			fmt.Println("channel full")
		}

	}
}

//判断管道中是否满了
func test06() {
	out:=make(chan string,10)
	go write(out)
	for data:=range out {
		fmt.Println("data==",data)
		time.Sleep(time.Millisecond*500)
	}
}

func b(ch chan string) {
	ch <- "bbb"
}

// 同时监听多个管道，直到其中一个管道ready
func test05() {
	out1 := make(chan string)
	out2 := make(chan string)
	go a(out1)
	go b(out2)
	select {
	case s1 := <-out1:
		fmt.Println("out1==", s1)
	case s2 := <-out2:
		fmt.Println("out2==", s2)
	//default:
	//	fmt.Println("执行默认的")
	}
}

func test04() {
	t1 := time.NewTicker(time.Second * 1)
	go func() {
		i := 0
		for {
			i++
			fmt.Println("t==", (<-t1.C).Format("2006-01-02 15:04:05"))
			if i == 5 {
				t1.Stop()
			}
		}
	}()
	//
	time.Sleep(time.Second * 6)
}

func test03() {
	t3 := time.NewTimer(time.Second * 2)
	<-t3.C
	fmt.Println("2秒到")

	<-time.After(time.Second * 3)
	fmt.Println("3秒到")
}

// 验证timer只能相应一次
func test02() {
	t1 := time.NewTimer(time.Second * 2)
	for {
		<-t1.C // 只执行第一次，第二次的时候会抛出死锁，因为通道里面没有值了
		fmt.Println("时间到")
	}
}

// timer的基本使用
func test01() {
	timer := time.NewTimer(time.Second * 2)
	t1 := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("t1==", t1)

	t2 := (<-timer.C).Format("2006-01-02 15:04:05")
	fmt.Println("t2==", t2)
}
