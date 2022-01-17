package main

import (
	"errors"
	"fmt"
	"golang.org/x/sync/singleflight"
	"sync"
)

var errorNotExist = errors.New("not exist")
var g = singleflight.Group{}

func main() {
	var wg sync.WaitGroup
	wg.Add(10)

	//模拟10个并发
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			//data, err := getData("key")
			data, err := getDataWithSingleFlight("key")
			if err != nil {
				fmt.Print(err)
				return
			}
			fmt.Println(data)
		}()
	}
	wg.Wait()
}

// getData 获取数据
func getData(key string) (string, error) {
	data, err := getDataFromCache(key)
	if err == errorNotExist {
		//模拟从db中获取数据
		data, err = getDataFromDB(key)
		if err != nil {
			fmt.Println(err)
			return "", err
		}

		//TOOD: set cache
	} else if err != nil {
		return "", err
	}
	return data, nil
}

// getDataWithSingleFlight 获取数据，区别在于若缓存中不存在对应的数据，统一时刻只允许一个请求到达数据库，
// 这个请求拿到数据后，其他请求不用访问数据库也可以拿到相同的数据了
func getDataWithSingleFlight(key string) (string, error) {
	data, err := getDataFromCache(key)
	if err == errorNotExist {
		//模拟从db中获取数据
		//data, err = getDataFromDB(key)
		value, err, _ := g.Do(key, func() (interface{}, error) {
			return getDataFromDB(key)
		})
		if err != nil {
			fmt.Println(err)
			return "", err
		}
		data = value.(string)

	} else if err != nil {
		return "", err
	}
	return data, nil
}

//模拟从cache中获取值，cache中无该值
func getDataFromCache(key string) (string, error) {
	return "", errorNotExist
}

// getDataFromDB 模拟从数据库中获取值
func getDataFromDB(key string) (string, error) {
	fmt.Printf("get %s from database", key)
	return "data", nil
}
