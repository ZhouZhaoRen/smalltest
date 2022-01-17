package main

import (
	"fmt"
	"github.com/robfig/cron"
	"sync"
	"time"
)

var datas chan int
var wg sync.WaitGroup

func init() {
	datas = make(chan int, 10)
}
func main() {
	go test01()
	go test02()
	time.Sleep(time.Second * 100)
}

func test02() {
	c := cron.New()
	spec := "*/1 * * * * ?"
	c.AddFunc(spec, func() {
		//for i := 0; i < 30; i++ {
			//fmt.Println(<-datas)
		//}

	})
	c.Start()
}

func test01() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			//time.Sleep(time.Millisecond * 10)
			datas <- i
			fmt.Println(len(datas))

		}(i)
	}
}
