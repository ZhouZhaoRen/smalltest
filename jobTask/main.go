package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Task struct {
	msg string   //任务的内容, 具体可以是一个复杂的结构体对象
	pri int   // 任务的优先级，在对同一个bucket的数据，可以按照优先级来处理
	idx int   //  bucket 的标识
	status bool // 任务标识，标识任务是否执行成功，是否需要删除
}

func (t *Task) runTask() {  //简单的执行任务
	fmt.Println("run message", t.msg)
	t.status = true
}

var taskList = map[int][]Task{}

func sendTask(idx int) {
	msg := fmt.Sprintf("task message %d", idx)
	pri := idx / 60
	idx = idx % 60

	task := Task{
		msg,
		pri,
		idx,
		false,
	}
	taskList[idx] = append(taskList[idx], task)
}

/**
 * 假设 i是任务的id号，表示有一个150个任务要进如队列审核
 */
func initTask() {
	for i := 0; i < 150; i++ {
		sendTask(i)
	}
}

var ticker = time.NewTicker(1 * time.Second)
var cc = 0 //轮片指针

func main() {
	c := make(chan os.Signal)
	status:=true
	signal.Notify(c,
		syscall.SIGKILL,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		os.Interrupt,
		os.Kill,
	)
	initTask()
	go func() {
		for {
			select {
			case <-ticker.C:
				for _, t := range taskList[cc] {
					if t.status == false {
						t.runTask()
					}
				}
				cc += 1
				cc = cc%60  //循环轮询
			case <-c: //监听 信号
				ticker.Stop()
				fmt.Println("kill task")
				status = false
				break

			}
		}
	}()
	for {// 常驻
		time.Sleep(1*time.Second)
		if status == false {
			break
		}
	}
}