package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

var datas chan int
var count chan bool

func init() {
	datas = make(chan int, 1000)
	count = make(chan bool, 1)
}

func main() {
	//for i := 0; i < 20000; i++ {
	//	wg.Add(1)
	//	go test03(i)
	//}
	//wg.Wait()
	//time.Sleep(time.Second * 2)
	////m.Range(func(key, value interface{}) bool {
	////	//cas.Add(1)
	////	fmt.Println(key)
	////	fmt.Println(value)
	////	return true
	////})
	//fmt.Println("cas==", cas.String())
	test06()
	//datas <- 2
	//fmt.Println(<-datas)
}

//
func test04() {
	//defer wg.Done()
	for i := 0; i < 20; i++ {
		//wg.Add(1)
		go func(i int) {
			for ; i < 50; i++ {
				datas <- i
			}
		}(i)
	}
}

func test06() {
	//wg.Add(1)
	go test04()
	c := cron.New()
	spec := "*/1 * * * * ?"
	c.AddFunc(spec, func() {
		select {
		case count <- true:
			defer func() {
				<-count
			}()
			time.Sleep(time.Second * 2)
			fmt.Println(getData(1))
		default:
			fmt.Println("还有任务在跑")
		}

	})
	c.AddFunc(spec, func() {
		select {
		case count <- true:
			defer func() {
				<-count
			}()
			time.Sleep(time.Second * 2)
			fmt.Println(getData(99999))
		default:
			fmt.Println("还有任务在跑")
		}
	})
	c.Start()
	//wg.Wait()
	time.Sleep(time.Hour * 1)
}

func getData(a int) []int {
	var result []int
	if a == 99999 && len(datas) < 500 {
		result = append(result, 99999)
	}

	timer := time.NewTimer(time.Second * 3)
	for i := 0; i < 100; i++ {
		select {
		case data, ok := <-datas:
			if !ok {
				continue
			}
			result = append(result, data)
		case <-timer.C:
			return result
		}
	}
	return result
}
