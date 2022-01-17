package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"math/rand"
	"time"
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
	now := time.Now()
	for i := 0; i < 10000; i++ {
		//isAllow("myCmd",i)
		isAllow2("cmd2")
	}
	fmt.Println("花费的时间=",time.Now().Sub(now).Milliseconds())
	fmt.Println("成功的次数=", success)
	fmt.Println("拒绝的次数=", fail)
}

var success = 0
var fail = 0

func isAllow2(cmd string) {
	script:="local k=KEYS[1]" +
		"local m=tonumber(redis.call('hget',k,'next') or 0)" +
		"local t=redis.call('time')" +
		"local n=tonumber(t[1])*1e6+tonumber(t[2])" +
		"if(m>n)then return m-n end " +
		"local c=tonumber(redis.call('hget',k,'cur') or 0)" +
		"if(n>m)then c=math.min(1000,c+(n-m)/1000) m=n end " +
		"local u=math.min(1,c)" +
		"redis.replicate_commands()" +
		"redis.call('hset',k,'cur',c-u)" +
		"redis.call('hset',k,'next',m+(1-u)*1000)" +
		"redis.call('expire',k,60)" +
		"return 0"
	res, err := redis.Int64(redisCli.Do("EVAL", script, 1, genKey(cmd)))
	if err != nil {
		fmt.Printf("err=%+v", err)
		return
	}
	//fmt.Println("res==", res)
	if res == 0 {
		success++
		fmt.Println("通过")
	} else {
		fail++
		fmt.Println("拒绝")
		return
	}

}

func isAllow(cmd string,i int) {
	rand.Seed(time.Now().UnixNano())
	member := fmt.Sprintf("%s_%d", "aa", i)
	now := getMilTime(time.Now())
	script := `local k=KEYS[1]
               local n=tonumber(ARGV[1])
               redis.call('ZREMRANGEBYSCORE',k,0,n-1000)
               local res=tonumber(redis.call('ZCOUNT',k,n-1000,n))
               if res>1000 then return -1 end 
               redis.call('ZADD',k,n,ARGV[2])
               return res+1
              `
	res, err := redis.Int64(redisCli.Do("EVAL", script, 1, genKey(cmd), now, member))
	if err != nil {
		fmt.Printf("err=%+v", err)
		return
	}
	fmt.Println("res==", res)
	if res > 0 {
		success++
		fmt.Println("通过")
	} else {
		fail++
		fmt.Println("拒绝")
		return
	}

}

func genKey(cmd string) string {
	return fmt.Sprintf("%s_%s", "crud", cmd)
}

func getMilTime(now time.Time) int64 {
	return now.UnixNano() / 1e6
}
