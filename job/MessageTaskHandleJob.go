package job

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	// 添加job
	JobContainerCache.AddJobAtOnce("*/10 * * * * ?", time.Minute, &MessageTaskHandleJob{
		Chans: make(chan bool, 1),
	})
}

// MessageTaskHandleJob 消息任务处理任务
type MessageTaskHandleJob struct {
	Chans chan bool // 控制任务同时存在数
}

// Run 运行
func (obj *MessageTaskHandleJob) Run() {
	select {
	case obj.Chans <- true:
		defer func() {
			<-obj.Chans
		}()

		// 随机一个延时数，避免多个机器同时执行
		delaySecond := rand.Intn(1000)
		time.Sleep(time.Duration(delaySecond) * time.Millisecond)

		// 读取任务


		// 处理
		for i := range []int{} {
			fmt.Println(i)
			obj.Handle()
		}

	default:
	}
}

func (obj *MessageTaskHandleJob) Handle() {


}
