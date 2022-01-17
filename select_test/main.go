package main

import (
	"fmt"
	"time"
)

func main() {
	test03()
}

// test01 程序中声明两个channel，分别为chan1和chan2，依次启动两个协程，分别向两个channel中写入一个数据就进入睡眠。
// select语句两个case分别检测chan1和chan2是否可读，如果都不可读则执行default语句。
func test01() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		chan1 <- 1
		time.Sleep(5 * time.Second)
	}()

	go func() {
		chan2 <- 1
		time.Sleep(5 * time.Second)
	}()

	select {
	case <-chan1:
		fmt.Println("chan1 ready.")
	case <-chan2:
		fmt.Println("chan2 ready.")
	default:
		fmt.Println("default")
	}

	fmt.Println("main exit.")
	// 最终执行结果是三种都有可能执行，因为启动的协程和select语句并不能保证顺序执行
}

// test02 程序中声明两个channel，分别为chan1和chan2，依次启动两个协程，协程会判断一个bool类型的变量writeFlag来决定是否要向channel中写入数据，
// 由于writeFlag永远为false，所以实际上协程什么也没做。select语句两个case分别检测chan1和chan2是否可读，这个select语句不包含default语句。
func test02() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	writeFlag := false
	go func() {
		for {
			if writeFlag {
				chan1 <- 1
			}
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			if writeFlag {
				chan2 <- 1
			}
			time.Sleep(time.Second)
		}
	}()

	select {
	case <-chan1:
		fmt.Println("chan1 ready.")
	case <-chan2:
		fmt.Println("chan2 ready.")
	}

	fmt.Println("main exit.")
	// select会按照随机的顺序检测各case语句中channel是否ready，如果某个case中的channel已经ready则执行相应的case语句然后退出select流程，
	// 如果所有的channel都未ready且没有default的话，则会阻塞等待各个channel。所以上述程序会一直阻塞。
}

// test03 程序中声明两个channel，分别为chan1和chan2，依次启动两个协程，协程分别关闭两个channel。select语句两个case分别检测chan1和chan2是否可读，这个select语句不包含default语句。
func test03() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		close(chan1)
	}()

	go func() {
		close(chan2)
	}()

	select {
	case <-chan1:
		fmt.Println("chan1 ready.")
	case <-chan2:
		fmt.Println("chan2 ready.")
	}

	fmt.Println("main exit.")
	// select会按照随机的顺序检测各case语句中channel是否ready，考虑到已关闭的channel也是可读的，所以上述程序中select不会阻塞，具体执行哪个case语句具是随机的。
}

// test04 上面程序中只有一个空的select语句。
func test04() {
	select {}
	// 对于空的select语句，程序会被阻塞，准确的说是当前协程被阻塞，同时Golang自带死锁检测机制，当发现当前协程再也没有机会被唤醒时，则会panic。所以上述程序会panic。
}
