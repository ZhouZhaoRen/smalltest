package main

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var redisCli redis.Conn

type Test struct {
	Id   int
	Name string
}

func main() {
	// 10.0.4.16
	var err error
	redisCli, err = redis.Dial("tcp", "1.117.76.139:6379", redis.DialPassword("123456"))
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	defer redisCli.Close()
	fmt.Println("链接成功")

	//test02()
	//test03()
	//test04()
	//test03()
	//test06()
	test05()
	//test07()
}

func test07() {
	data, err := redis.Bytes(redisCli.Do("getrange", "t2", 8, -1))
	if err != nil {
		fmt.Println(err)
		return
	}
	var test Test
	err = json.Unmarshal(data, &test)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(test)
}

func test06() {
	test := Test{
		Id:   1,
		Name: "zzr",
	}
	data, _ := json.Marshal(test)
	script := `
	local newCas=tonumber(ARGV[2]) --传入版本号
	local packValue = struct.pack('<I8c0', newCas, ARGV[1]);  --版本号和值组装起来
    redis.call('set', KEYS[1], packValue);     --set value
`
	_, err := redisCli.Do("eval", script, 1, "t2", data, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func test05() {

	script := `
	-- 输入：参数2， KEYS[1]=目标Key， ARGV[1] =当前对应的值(当前Cas+1 + value )，ARGV[2]= 当前cas
local key     =KEYS[1];
local newValue=ARGV[1];
local checkCAS=tonumber(ARGV[2]);
local oldRaw  =redis.call('getrange', key, 0, 7);
local newCAS  = ( checkCAS >= 4294967295 ) and 0 or ( checkCAS + 1 )
local packValue = struct.pack('<I8c0', newCAS, newValue);  --版本号和值组装起来
-- 数据不存在直接覆盖
if oldRaw == "" then
	redis.call('set', key, packValue);
	return {0, newCAS, 1 } ;
end

local oldCAS = struct.unpack('<I8', oldRaw); --这里是将value中头8字节解析成无符号整数，即数据版本号
-- cas不匹配
if oldCAS ~= checkCAS then
	return {4, oldCAS, 0 } ;
end
redis.call('set', key, packValue);
return {0, newCAS, 2 } ;

-- 请求数据 checkCAS 是上一次拉取的到的版本号， newvalue的前8字节 是  checkCAS+1
`
	test := Test{
		Id:   4,
		Name: "zzr3",
	}
	data, _ := json.Marshal(test)
	res, err := redis.Values(redisCli.Do("eval", script, 1, "t3", data, 2))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func test04() {

	data, _ := json.Marshal(1)
	test := Test{Id: 1, Name: "zzr"}
	data2, _ := json.Marshal(test)
	fmt.Println(data)
	fmt.Println(data2)
	data = append(data, data2...)
	fmt.Println()
	redisCli.Do("set", "t1", data)
}

func test03() {
	data, err := redis.Bytes(redisCli.Do("get", "t1"))
	if err != nil {
		fmt.Println("test03:", err)
		return
	}
	fmt.Println(data)
	var a int
	err = json.Unmarshal(data, &a)
	if err != nil {
		fmt.Println("Unmarshal:", err)
		return
	}
	fmt.Println(a)
}

func test02() {
	script := "local m=tonumber(redis.call('GET',KEYS[1]) or 0) if (m>tonumber(ARGV[2])) then return -1 end if (m<tonumber(ARGV[1])) then return -2 end"
	//script:="local m=tonumber(redis.call('GET',KEYS[1]) or 0) return m "
	data, err := redis.Int64(redisCli.Do("eval", script, 1, "small"))
	if err != nil {
		fmt.Println("err==", err)
		return
	}
	fmt.Println(data)
}

func test01() {
	script := `local function test(val)
    local ret1 = {1, 2}
    local ret2 = "hello"
    local ret3 = val
    local ret = {}
    ret[1] = ret1
    ret[2] = ret2
    ret[3] = ret3
    return ret
end
return test(KEYS[1])
`

	data, err := redis.Strings(redisCli.Do("EVAL", script, 1, 3))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}
