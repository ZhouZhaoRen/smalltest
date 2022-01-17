package job

import (
	"fmt"
	"time"
)

// DefaultJobContainer 默认的jobContainer
var JobContainerCache JobContainer

// JobContainer job容器：承接所有job信息
type JobContainer struct {
	jobs []*JobInfo // job数组
}

// JobInfo job信息
type JobInfo struct {
	spec    string        // 任务执行规格，如："0 */5 * * * ?" 每隔5分钟
	timeout time.Duration // 任务执行超时时间
	job     Job           // 任务
	atOnce  bool          // 启动时立即执行一次
}

// AddJob 往容器中添加job
func (obj *JobContainer) AddJob(spec string, timeout time.Duration, job Job) {
	obj.jobs = append(obj.jobs, &JobInfo{spec, timeout, job, false})
}

// AddJobAtOnce 往容器中添加job，且启动时立即执行一次
func (obj *JobContainer) AddJobAtOnce(spec string, timeout time.Duration, job Job) {
	obj.jobs = append(obj.jobs, &JobInfo{spec, timeout, job, true})
}

// Run 运行job容器
func (obj *JobContainer) Run() {
	for i := range obj.jobs {
		if obj.jobs[i].atOnce {
			fmt.Println("立马执行任务:",obj.jobs[i])
		} else {
			time.Sleep(time.Second*3)
			fmt.Println("延迟三秒后执行任务:",obj.jobs[i])
			//cat.StartTask(obj.jobs[i].spec, obj.jobs[i].timeout, obj.jobs[i].job.Run)
		}

	}
}