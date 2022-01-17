package main

import "fmt"

// worker 当select内部随机函数得到的是执行ch2的时候，先检查ch1是否就绪，否则再执行ch2
func worker(ch1, ch2 chan int, stop chan struct{}) {
	for {
		select {
		case <-stop:
			return
		case job1 := <-ch1:
			fmt.Println(job1)
		case job2 := <-ch2:
		priority:
			// 当ch1就绪时就会一直执行ch1
			for {
				select {
				case job1 := <-ch1:
					fmt.Println(job1)
				default:
					break priority
				}
			}
			fmt.Println(job2)
		}
	}
}
