package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var (
	redisCli redis.Conn
)

func main() {
	//第一种连接方法
	var err error
	redisCli, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	defer redisCli.Close()
	test05()
}

// 返回数组，由于LUA脚本只能返回一个值，可以用大括号括起来
func test05() {
	script := `local key=tonumber(redis.call('get',KEYS[1]))
local value=tonumber(ARGV[1])
if(key==value) then return {key,value}
else return {-1,-1} end
`
	data, err := redis.Int64s(redisCli.Do("EVAL", script, 1, "hh", 2))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}

// 流程控制
func test04() {
	script := `local key=KEYS[1]
local value=ARGV[1]
if(key==value) then return key
else return value end
`
	data, err := redis.String(redisCli.Do("EVAL", script, 1, "key", "value"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}

// 最简单的操作，变量和返回值
func test03() {
	script := `local key=KEYS[1]
local value=ARGV[1]
return redis.call('SET',key,value)
`
	data, err := redisCli.Do("EVAL", script, 1, "key", "value")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}

func test02() {
	res, err := redis.Values(redisCli.Do("hscan", "test", 0))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	result, err := redis.Int64Map(res[1], err)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
	//for key, value := range res {
	//	fmt.Printf("key==%s  value==%d", key, value)
	//}

}

func test01() {
	//
	redisCli.Do("set", "name", "small")
	res, err := redis.String(redisCli.Do("get", "name"))
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println(res)
}
