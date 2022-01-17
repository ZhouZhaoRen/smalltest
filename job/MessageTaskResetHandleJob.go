package job

import (
	"math/rand"
	"time"
)

func init() {
	// 添加job
	JobContainerCache.AddJobAtOnce("0 */1 * * * ?", time.Minute, &MessageTaskResetHandleJob{
		chans: make(chan bool, 1),
	})
}

// MessageTaskResetHandleJob 消息任务重置处理任务
type MessageTaskResetHandleJob struct {
	chans chan bool // 控制任务同时存在数
}

// Run 运行
func (obj *MessageTaskResetHandleJob) Run() {
	select {
	case obj.chans <- true:
		defer func() {
			<-obj.chans
		}()

		// 随机一个延时数，避免多个机器同时执行
		delaySecond := rand.Intn(10)
		time.Sleep(time.Duration(delaySecond) * time.Second)

		// 更新记录为初始状态

	default:
	}
}